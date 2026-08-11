package win

import (
	"fmt"

	"github.com/yukumo-group/yukumo-script/cmd/cmdinterface"
	"github.com/yukumo-group/yukumo-script/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/language"
	"github.com/yukumo-group/yukumo-script/phontsmanager"
	"github.com/yukumo-group/yukumo-script/utils"
	"github.com/yukumo-group/yukumo-script/utils/audio"
	"github.com/yukumo-group/yukumo-script/utils/osoperation"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
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
	Short: "generation command contains subcommands that can generate audios through pre-defined characters or phont files",
	Long: `
subcommands of generation command allow you to generate your audios
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
		title.Println("Here are the available phonts:")
		for _, phontName := range phontsmanager.PhontNameToFileName.GetAllKeys() {
			text.Println(phontName)
		}
		title.Println("Input the name of the phont you want to use to generate audio:")
		// Input
		var phontName string
		_, errInput := fmt.Scan(&phontName)
		if errInput != nil {
			cmdLogger.Error(errInput.Error())
			errMessage.Println(errInput.Error())
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
			errMessage.Printf(
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
			errMessage.Printf(
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
			errMessage.Println(errConvertLang.Error())
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
			errMessage.Println(errCreateTask.Error())
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
			errMessage.Printf(
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
		errGenerate := newTask.GenerateWin(
			phontPath,
			utils.ResultDir,
		)
		if errGenerate != nil {
			cmdLogger.Error(errGenerate.Error())
			errMessage.Println(errGenerate.Error())
			return
		}
		resultFile, errSaveFile := newTask.SaveFile(
			utils.SingleSentenceDir,
		)
		if errSaveFile != nil {
			cmdLogger.Error(errSaveFile.Error())
			errMessage.Println(errSaveFile.Error())
			return
		}
		errNewTask := singlesentence.Manager.NewTask(
			newTask.TaskName,
			resultFile,
		)
		if errNewTask != nil {
			cmdLogger.Error(errNewTask.Error())
			errMessage.Println(errNewTask.Error())
			return
		}
		// Show file info
		title.Printf(
			"Generated File Path: %s\n",
			*newTask.ResultFile,
		)
		// Export file
		if newTask.ResultFile == nil {
			cmdLogger.Error("The directory of the result file cannot be nil")
			errMessage.Println("The directory of the result file cannot be nil")
			return
		}
		doExport, errAskExport := cmdinterface.YesOrNoWithColor(
			title,
			"Do you want to export the file generated? (Warning: This will overwrite the file if the path of the exported file already exists)",
			false,
		)
		if errAskExport != nil {
			cmdLogger.Error(errAskExport.Error())
			errMessage.Println(errAskExport.Error())
			return
		}
		if !doExport {
			return
		}
		title.Println("Input the directory of the exported file to store")
		var exportDirectory string
		fmt.Scan(&exportDirectory)
		exportDirectory = osoperation.ParseWindowsPath(exportDirectory)
		title.Println("Input the name of the file to store (no suffix needed)")
		var exportFileName string
		fmt.Scan(&exportFileName)
		doChangeFormat, errAskChangeFormat := cmdinterface.YesOrNoWithColor(
			title,
			"Do you want to change the format of the exported file to formats other than wav file?",
			false,
		)
		if errAskChangeFormat != nil {
			cmdLogger.Error(errAskChangeFormat.Error())
			errMessage.Println(errAskChangeFormat.Error())
			return
		}
		if doChangeFormat {
			title.Println("Select the format of the exported file")
			formats := audio.GetAllFormats()
			for _, format := range formats {
				text.Println(format.ToString())
			}
			var selectedFormat string
			fmt.Scan(&selectedFormat)
			errConvertAudioFormat := audio.ConvertAll(
				*newTask.ResultFile,
				exportDirectory,
				exportFileName,
				audio.ToFormat(selectedFormat),
			)
			if errConvertAudioFormat != nil {
				cmdLogger.Error(errConvertAudioFormat.Error())
				errMessage.Println(errConvertAudioFormat.Error())
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
			errMessage.Println(errCopyFile.Error())
			return
		}
	},
}
