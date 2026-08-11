package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/yukumo-group/yukumo-script/internal/cmdinterface"
	"github.com/yukumo-group/yukumo-script/pkg/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/pkg/language"
	"github.com/yukumo-group/yukumo-script/pkg/phontsmanager"
	"github.com/yukumo-group/yukumo-script/pkg/utils"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
	"github.com/yukumo-group/yukumo-script/pkg/utils/osoperation"
)

// SingleSentenceTaskNameByFile defines the name of the task
var SingleSentenceTaskNameByFile string

// SingleSentenceTaskTextByFile defines the text to generate for the task
var SingleSentenceTaskTextByFile string

// SingleSentenceTaskSpeedByFile defines the speed of the audio
var SingleSentenceTaskSpeedByFile int

// SingleSentenceTaskLanguageByFile defines the language of the audio
var SingleSentenceTaskLanguageByFile int

// generationCommand contains commands for generation
var generationCMD = &cobra.Command{
	Use:   "generation",
	Short: "generation command contains subcommands that can generate audio through pre-defined characters or phont files",
	Long: `
subcommands of generation command allow you to generate your audio
	`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// generateByFile generates wav by file
var generateByFileCMD = &cobra.Command{
	Use:   "generateByFile",
	Short: "Generate yukumo audio through the file",
	Long: `
generateByFile allows you to generate yukumo audio through phont file directly
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Define the format of the texts
		title := color.New(color.FgGreen).Add(color.Bold)
		errMessage := color.New(color.FgRed).Add(color.Bold)
		text := color.New(color.Italic)
		// Print info
		_, _ = title.Println("Here are the available phonts:")
		for _, phontName := range phontsmanager.PhontNameToFileName.GetAllKeys() {
			_, _ = text.Println(phontName)
		}
		_, _ = title.Println("Input the name of the phont you want to use to generate audio:")
		// Input
		var phontName string
		_, errInput := fmt.Scan(&phontName)
		if errInput != nil {
			cmdLogger.Error(errInput.Error())
			_, _ = errMessage.Println(errInput.Error())
			return
		}
		// Get the phont
		_, phontExists := phontsmanager.PhontNameToFileName.GetValue(
			phontName,
		)
		if !phontExists {
			cmdLogger.Error(
				fmt.Sprintf(
					"No such phont %s",
					phontName,
				),
			)
			_, _ = errMessage.Printf(
				"No such phont %s",
				phontName,
			)
			return
		}
		if singlesentence.Manager.HasTask(SingleSentenceTaskNameByFile) {
			cmdLogger.Error(
				fmt.Sprintf(
					"task %s already exists",
					SingleSentenceTaskNameByFile,
				),
			)
			_, _ = errMessage.Printf(
				"task %s already exists",
				SingleSentenceTaskNameByFile,
			)
			return
		}
		// Create New Task
		processedText, errConvertLang := language.ConvertText(
			SingleSentenceTaskTextByFile,
			language.ToLanguage(SingleSentenceTaskLanguageByFile),
		)
		if errConvertLang != nil {
			cmdLogger.Error(errConvertLang.Error())
			_, _ = errMessage.Println(errConvertLang.Error())
			return
		}
		newTask, errCreateTask := singlesentence.NewSingleSentenceTask(
			processedText,
			nil,
			&phontName,
			SingleSentenceTaskSpeedByFile,
			SingleSentenceTaskNameByFile,
		)
		if errCreateTask != nil {
			cmdLogger.Error(errCreateTask.Error())
			_, _ = errMessage.Println(errCreateTask.Error())
			return
		}
		// Generate
		phontFileName, phontFileExists := phontsmanager.PhontNameToFileName.GetValue(
			*newTask.PhontName,
		)
		if !phontFileExists {
			cmdLogger.Error(
				fmt.Sprintf(
					"No such phont file %s",
					phontName,
				),
			)
			_, _ = errMessage.Printf(
				"No such phont file %s",
				phontName,
			)
			return
		}
		phontPath := fmt.Sprintf(
			"%s/%s",
			utils.PhontsDir,
			phontFileName,
		)
		errGenerate := newTask.Generate(
			phontPath,
			utils.ResultDir,
		)
		if errGenerate != nil {
			cmdLogger.Error(errGenerate.Error())
			_, _ = errMessage.Println(errGenerate.Error())
			return
		}
		resultFile, errSaveFile := newTask.SaveFile(
			utils.SingleSentenceDir,
		)
		if errSaveFile != nil {
			cmdLogger.Error(errSaveFile.Error())
			_, _ = errMessage.Println(errSaveFile.Error())
			return
		}
		errNewTask := singlesentence.Manager.NewTask(
			newTask.TaskName,
			resultFile,
		)
		if errNewTask != nil {
			cmdLogger.Error(errNewTask.Error())
			_, _ = errMessage.Println(errNewTask.Error())
			return
		}
		// Show file info
		_, _ = title.Printf(
			"Generated File Path: %s\n",
			*newTask.ResultFile,
		)
		// Export file
		if newTask.ResultFile == nil {
			cmdLogger.Error("The directory of the result file cannot be nil")
			_, _ = errMessage.Println("The directory of the result file cannot be nil")
			return
		}
		doExport, errAskExport := cmdinterface.YesOrNoWithColor(
			title,
			"Do you want to export the file generated? (Warning: This will overwrite the file if the path of the exported file already exists)",
			false,
		)
		if errAskExport != nil {
			cmdLogger.Error(errAskExport.Error())
			_, _ = errMessage.Println(errAskExport.Error())
			return
		}
		if !doExport {
			return
		}
		_, _ = title.Println("Input the directory of the exported file to store")
		var exportDirectory string
		if _, err := fmt.Scan(&exportDirectory); err != nil {
			cmdLogger.Error(err.Error())
			_, _ = errMessage.Println(err.Error())
			return
		}
		exportDirectory = osoperation.ParseWindowsPath(exportDirectory)
		_, _ = title.Println("Input the name of the file to store (no suffix needed)")
		var exportFileName string
		if _, err := fmt.Scan(&exportFileName); err != nil {
			cmdLogger.Error(err.Error())
			_, _ = errMessage.Println(err.Error())
			return
		}
		doChangeFormat, errAskChangeFormat := cmdinterface.YesOrNoWithColor(
			title,
			"Do you want to change the format of the exported file to formats other than wav file?",
			false,
		)
		if errAskChangeFormat != nil {
			cmdLogger.Error(errAskChangeFormat.Error())
			_, _ = errMessage.Println(errAskChangeFormat.Error())
			return
		}
		if doChangeFormat {
			_, _ = title.Println("Select the format of the exported file")
			formats := audio.GetAllFormats()
			for _, format := range formats {
				_, _ = text.Println(format.ToString())
			}
			var selectedFormat string
			if _, err := fmt.Scan(&selectedFormat); err != nil {
				cmdLogger.Error(err.Error())
				_, _ = errMessage.Println(err.Error())
				return
			}
			errConvertAudioFormat := audio.ConvertAll(
				*newTask.ResultFile,
				exportDirectory,
				exportFileName,
				audio.ToFormat(selectedFormat),
			)
			if errConvertAudioFormat != nil {
				cmdLogger.Error(errConvertAudioFormat.Error())
				_, _ = errMessage.Println(errConvertAudioFormat.Error())
				return
			}
			return
		}
		errCopyFile := osoperation.CopyFile(
			*newTask.ResultFile,
			exportDirectory,
			exportFileName,
			"wav",
		)
		if errCopyFile != nil {
			cmdLogger.Error(errCopyFile.Error())
			_, _ = errMessage.Println(errCopyFile.Error())
			return
		}
	},
}
