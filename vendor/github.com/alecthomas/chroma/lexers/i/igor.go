package i

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Igor lexer.
var Igor = internal.Register(MustNewLexer(
	&Config{
		Name:            "Igor",
		Aliases:         []string{"igor", "igorpro"},
		Filenames:       []string{"*.ipf"},
		MimeTypes:       []string{"text/ipf"},
		CaseInsensitive: true,
	},
	Rules{
		"root": {
			{`//.*$`, CommentSingle, nil},
			{`"([^"\\]|\\.)*"`, LiteralString, nil},
			{Words(`\b`, `\b`, `if`, `else`, `elseif`, `endif`, `for`, `endfor`, `strswitch`, `switch`, `case`, `default`, `endswitch`, `do`, `while`, `try`, `catch`, `endtry`, `break`, `continue`, `return`, `AbortOnRTE`, `AbortOnValue`), Keyword, nil},
			{Words(`\b`, `\b`, `variable`, `string`, `constant`, `strconstant`, `NVAR`, `SVAR`, `WAVE`, `STRUCT`, `dfref`, `funcref`, `char`, `uchar`, `int16`, `uint16`, `int32`, `uint32`, `int64`, `uint64`, `float`, `double`), KeywordType, nil},
			{Words(`\b`, `\b`, `override`, `ThreadSafe`, `MultiThread`, `static`, `Proc`, `Picture`, `Prompt`, `DoPrompt`, `macro`, `window`, `function`, `end`, `Structure`, `EndStructure`, `EndMacro`, `Menu`, `SubMenu`), KeywordReserved, nil},
			{Words(`\b`, `\b`, `Abort`, `AddFIFOData`, `AddFIFOVectData`, `AddMovieAudio`, `AddMovieFrame`, `AddWavesToBoxPlot`, `AddWavesToViolinPlot`, `AdoptFiles`, `APMath`, `Append`, `AppendBoxPlot`, `AppendImage`, `AppendLayoutObject`, `AppendMatrixContour`, `AppendText`, `AppendToGizmo`, `AppendToGraph`, `AppendToLayout`, `AppendToTable`, `AppendViolinPlot`, `AppendXYZContour`, `AutoPositionWindow`, `AxonTelegraphFindServers`, `BackgroundInfo`, `Beep`, `BoundingBall`, `BoxSmooth`, `BrowseURL`, `BuildMenu`, `Button`, `cd`, `Chart`, `CheckBox`, `CheckDisplayed`, `ChooseColor`, `Close`, `CloseHelp`, `CloseMovie`, `CloseProc`, `ColorScale`, `ColorTab2Wave`, `Concatenate`, `ControlBar`, `ControlInfo`, `ControlUpdate`, `ConvertGlobalStringTextEncoding`, `ConvexHull`, `Convolve`, `CopyDimLabels`, `CopyFile`, `CopyFolder`, `CopyScales`, `Correlate`, `CreateAliasShortcut`, `CreateBrowser`, `Cross`, `CtrlBackground`, `CtrlFIFO`, `CtrlNamedBackground`, `Cursor`, `CurveFit`, `CustomControl`, `CWT`, `DAQmx_AI_SetupReader`, `DAQmx_AO_SetOutputs`, `DAQmx_CTR_CountEdges`, `DAQmx_CTR_OutputPulse`, `DAQmx_CTR_Period`, `DAQmx_CTR_PulseWidth`, `DAQmx_DIO_Config`, `DAQmx_DIO_WriteNewData`, `DAQmx_Scan`, `DAQmx_WaveformGen`, `Debugger`, `DebuggerOptions`, `DefaultFont`, `DefaultGuiControls`, `DefaultGuiFont`, `DefaultTextEncoding`, `DefineGuide`, `DelayUpdate`, `DeleteAnnotations`, `DeleteFile`, `DeleteFolder`, `DeletePoints`, `Differentiate`, `dir`, `Display`, `DisplayHelpTopic`, `DisplayProcedure`, `DoAlert`, `DoIgorMenu`, `DoUpdate`, `DoWindow`, `DoXOPIdle`, `DPSS`, `DrawAction`, `DrawArc`, `DrawBezier`, `DrawLine`, `DrawOval`, `DrawPICT`, `DrawPoly`, `DrawRect`, `DrawRRect`, `DrawText`, `DrawUserShape`, `DSPDetrend`, `DSPPeriodogram`, `Duplicate`, `DuplicateDataFolder`, `DWT`, `EdgeStats`, `Edit`, `ErrorBars`, `EstimatePeakSizes`, `Execute`, `ExecuteScriptText`, `ExperimentInfo`, `ExperimentModified`, `ExportGizmo`, `Extract`, `FastGaussTransform`, `FastOp`, `FBinRead`, `FBinWrite`, `FFT`, `FGetPos`, `FIFOStatus`, `FIFO2Wave`, `FilterFIR`, `FilterIIR`, `FindAPeak`, `FindContour`, `FindDuplicates`, `FindLevel`, `FindLevels`, `FindPeak`, `FindPointsInPoly`, `FindRoots`, `FindSequence`, `FindValue`, `FMaxFlat`, `FPClustering`, `fprintf`, `FReadLine`, `FSetPos`, `FStatus`, `FTPCreateDirectory`, `FTPDelete`, `FTPDownload`, `FTPUpload`, `FuncFit`, `FuncFitMD`, `GBLoadWave`, `GetAxis`, `GetCamera`, `GetFileFolderInfo`, `GetGizmo`, `GetLastUserMenuInfo`, `GetMarquee`, `GetMouse`, `GetSelection`, `GetWindow`, `GISCreateVectorLayer`, `GISGetRasterInfo`, `GISGetRegisteredFileInfo`, `GISGetVectorLayerInfo`, `GISLoadRasterData`, `GISLoadVectorData`, `GISRasterizeVectorData`, `GISRegisterFile`, `GISTransformCoords`, `GISUnRegisterFile`, `GISWriteFieldData`, `GISWriteGeometryData`, `GISWriteRaster`, `GPIBReadBinaryWave2`, `GPIBReadBinary2`, `GPIBReadWave2`, `GPIBRead2`, `GPIBWriteBinaryWave2`, `GPIBWriteBinary2`, `GPIBWriteWave2`, `GPIBWrite2`, `GPIB2`, `GraphNormal`, `GraphWaveDraw`, `GraphWaveEdit`, `Grep`, `GroupBox`, `Hanning`, `HDFInfo`, `HDFReadImage`, `HDFReadSDS`, `HDFReadVset`, `HDF5CloseFile`, `HDF5CloseGroup`, `HDF5ConvertColors`, `HDF5CreateFile`, `HDF5CreateGroup`, `HDF5CreateLink`, `HDF5Dump`, `HDF5DumpErrors`, `HDF5DumpState`, `HDF5FlushFile`, `HDF5ListAttributes`, `HDF5ListGroup`, `HDF5LoadData`, `HDF5LoadGroup`, `HDF5LoadImage`, `HDF5OpenFile`, `HDF5OpenGroup`, `HDF5SaveData`, `HDF5SaveGroup`, `HDF5SaveImage`, `HDF5TestOperation`, `HDF5UnlinkObject`, `HideIgorMenus`, `HideInfo`, `HideProcedures`, `HideTools`, `HilbertTransform`, `Histogram`, `ICA`, `IFFT`, `ImageAnalyzeParticles`, `ImageBlend`, `ImageBoundaryToMask`, `ImageComposite`, `ImageEdgeDetection`, `ImageFileInfo`, `ImageFilter`, `ImageFocus`, `ImageFromXYZ`, `ImageGenerateROIMask`, `ImageGLCM`, `ImageHistModification`, `ImageHistogram`, `ImageInterpolate`, `ImageLineProfile`, `ImageLoad`, `ImageMorphology`, `ImageRegistration`, `ImageRemoveBackground`, `ImageRestore`, `ImageRotate`, `ImageSave`, `ImageSeedFill`, `ImageSkeleton3d`, `ImageSnake`, `ImageStats`, `ImageThreshold`, `ImageTransform`, `ImageUnwrapPhase`, `ImageWindow`, `IndexSort`, `InsertPoints`, `Integrate`, `IntegrateODE`, `Integrate2D`, `Interpolate2`, `Interpolate3D`, `Interp3DPath`, `ITCCloseAll2`, `ITCCloseDevice2`, `ITCConfigAllChannels2`, `ITCConfigChannelReset2`, `ITCConfigChannelUpload2`, `ITCConfigChannel2`, `ITCFIFOAvailableAll2`, `ITCFIFOAvailable2`, `ITCGetAllChannelsConfig2`, `ITCGetChannelConfig2`, `ITCGetCurrentDevice2`, `ITCGetDeviceInfo2`, `ITCGetDevices2`, `ITCGetErrorString2`, `ITCGetSerialNumber2`, `ITCGetState2`, `ITCGetVersions2`, `ITCInitialize2`, `ITCOpenDevice2`, `ITCReadADC2`, `ITCReadDigital2`, `ITCReadTimer2`, `ITCSelectDevice2`, `ITCSetDAC2`, `ITCSetGlobals2`, `ITCSetModes2`, `ITCSetState2`, `ITCStartAcq2`, `ITCStopAcq2`, `ITCUpdateFIFOPositionAll2`, `ITCUpdateFIFOPosition2`, `ITCWriteDigital2`, `JCAMPLoadWave`, `JointHistogram`, `KillBackground`, `KillControl`, `KillDataFolder`, `KillFIFO`, `KillFreeAxis`, `KillPath`, `KillPICTs`, `KillStrings`, `KillVariables`, `KillWaves`, `KillWindow`, `KMeans`, `Label`, `Layout`, `LayoutPageAction`, `LayoutSlideShow`, `Legend`, `LinearFeedbackShiftRegister`, `ListBox`, `LoadData`, `LoadPackagePreferences`, `LoadPICT`, `LoadWave`, `Loess`, `LombPeriodogram`, `Make`, `MakeIndex`, `MarkPerfTestTime`, `MatrixConvolve`, `MatrixCorr`, `MatrixEigenV`, `MatrixFilter`, `MatrixGaussJ`, `MatrixGLM`, `MatrixInverse`, `MatrixLinearSolve`, `MatrixLinearSolveTD`, `MatrixLLS`, `MatrixLUBkSub`, `MatrixLUD`, `MatrixLUDTD`, `MatrixMultiply`, `MatrixOP`, `MatrixSchur`, `MatrixSolve`, `MatrixSVBkSub`, `MatrixSVD`, `MatrixTranspose`, `MCC_FindServers`, `MeasureStyledText`, `MFR_CheckForNewBricklets`, `MFR_CloseResultFile`, `MFR_CreateOverviewTable`, `MFR_GetBrickletCount`, `MFR_GetBrickletData`, `MFR_GetBrickletDeployData`, `MFR_GetBrickletMetaData`, `MFR_GetBrickletRawData`, `MFR_GetReportTemplate`, `MFR_GetResultFileMetaData`, `MFR_GetResultFileName`, `MFR_GetVernissageVersion`, `MFR_GetVersion`, `MFR_GetXOPErrorMessage`, `MFR_OpenResultFile`, `MLLoadWave`, `Modify`, `ModifyBoxPlot`, `ModifyBrowser`, `ModifyCamera`, `ModifyContour`, `ModifyControl`, `ModifyControlList`, `ModifyFreeAxis`, `ModifyGizmo`, `ModifyGraph`, `ModifyImage`, `ModifyLayout`, `ModifyPanel`, `ModifyTable`, `ModifyViolinPlot`, `ModifyWaterfall`, `MoveDataFolder`, `MoveFile`, `MoveFolder`, `MoveString`, `MoveSubwindow`, `MoveVariable`, `MoveWave`, `MoveWindow`, `MultiTaperPSD`, `MultiThreadingControl`, `NC_CloseFile`, `NC_DumpErrors`, `NC_Inquire`, `NC_ListAttributes`, `NC_ListObjects`, `NC_LoadData`, `NC_OpenFile`, `NeuralNetworkRun`, `NeuralNetworkTrain`, `NewCamera`, `NewDataFolder`, `NewFIFO`, `NewFIFOChan`, `NewFreeAxis`, `NewGizmo`, `NewImage`, `NewLayout`, `NewMovie`, `NewNotebook`, `NewPanel`, `NewPath`, `NewWaterfall`, `NILoadWave`, `NI4882`, `Note`, `Notebook`, `NotebookAction`, `Open`, `OpenHelp`, `OpenNotebook`, `Optimize`, `ParseOperationTemplate`, `PathInfo`, `PauseForUser`, `PauseUpdate`, `PCA`, `PlayMovie`, `PlayMovieAction`, `PlaySound`, `PopupContextualMenu`, `PopupMenu`, `Preferences`, `PrimeFactors`, `Print`, `printf`, `PrintGraphs`, `PrintLayout`, `PrintNotebook`, `PrintSettings`, `PrintTable`, `Project`, `PulseStats`, `PutScrapText`, `pwd`, `Quit`, `RatioFromNumber`, `Redimension`, `Remez`, `Remove`, `RemoveContour`, `RemoveFromGizmo`, `RemoveFromGraph`, `RemoveFromLayout`, `RemoveFromTable`, `RemoveImage`, `RemoveLayoutObjects`, `RemovePath`, `Rename`, `RenameDataFolder`, `RenamePath`, `RenamePICT`, `RenameWindow`, `ReorderImages`, `ReorderTraces`, `ReplaceText`, `ReplaceWave`, `Resample`, `ResumeUpdate`, `Reverse`, `Rotate`, `Save`, `SaveData`, `SaveExperiment`, `SaveGizmoCopy`, `SaveGraphCopy`, `SaveNotebook`, `SavePackagePreferences`, `SavePICT`, `SaveTableCopy`, `SetActiveSubwindow`, `SetAxis`, `SetBackground`, `SetDashPattern`, `SetDataFolder`, `SetDimLabel`, `SetDrawEnv`, `SetDrawLayer`, `SetFileFolderInfo`, `SetFormula`, `SetIdlePeriod`, `SetIgorHook`, `SetIgorMenuMode`, `SetIgorOption`, `SetMarquee`, `SetProcessSleep`, `SetRandomSeed`, `SetScale`, `SetVariable`, `SetWaveLock`, `SetWaveTextEncoding`, `SetWindow`, `ShowIgorMenus`, `ShowInfo`, `ShowTools`, `Silent`, `Sleep`, `Slider`, `Smooth`, `SmoothCustom`, `Sort`, `SortColumns`, `SoundInRecord`, `SoundInSet`, `SoundInStartChart`, `SoundInStatus`, `SoundInStopChart`, `SoundLoadWave`, `SoundSaveWave`, `SphericalInterpolate`, `SphericalTriangulate`, `SplitString`, `SplitWave`, `sprintf`, `SQLHighLevelOp`, `sscanf`, `Stack`, `StackWindows`, `StatsAngularDistanceTest`, `StatsANOVA1Test`, `StatsANOVA2NRTest`, `StatsANOVA2RMTest`, `StatsANOVA2Test`, `StatsChiTest`, `StatsCircularCorrelationTest`, `StatsCircularMeans`, `StatsCircularMoments`, `StatsCircularTwoSampleTest`, `StatsCochranTest`, `StatsContingencyTable`, `StatsDIPTest`, `StatsDunnettTest`, `StatsFriedmanTest`, `StatsFTest`, `StatsHodgesAjneTest`, `StatsJBTest`, `StatsKDE`, `StatsKendallTauTest`, `StatsKSTest`, `StatsKWTest`, `StatsLinearCorrelationTest`, `StatsLinearRegression`, `StatsMultiCorrelationTest`, `StatsNPMCTest`, `StatsNPNominalSRTest`, `StatsQuantiles`, `StatsRankCorrelationTest`, `StatsResample`, `StatsSample`, `StatsScheffeTest`, `StatsShapiroWilkTest`, `StatsSignTest`, `StatsSRTest`, `StatsTTest`, `StatsTukeyTest`, `StatsVariancesTest`, `StatsWatsonUSquaredTest`, `StatsWatsonWilliamsTest`, `StatsWheelerWatsonTest`, `StatsWilcoxonRankTest`, `StatsWRCorrelationTest`, `STFT`, `String`, `StructFill`, `StructGet`, `StructPut`, `SumDimension`, `SumSeries`, `TabControl`, `Tag`, `TDMLoadData`, `TDMSaveData`, `TextBox`, `ThreadGroupPutDF`, `ThreadStart`, `TickWavesFromAxis`, `Tile`, `TileWindows`, `TitleBox`, `ToCommandLine`, `ToolsGrid`, `Triangulate3d`, `Unwrap`, `URLRequest`, `ValDisplay`, `Variable`, `VDTClosePort2`, `VDTGetPortList2`, `VDTGetStatus2`, `VDTOpenPort2`, `VDTOperationsPort2`, `VDTReadBinaryWave2`, `VDTReadBinary2`, `VDTReadHexWave2`, `VDTReadHex2`, `VDTReadWave2`, `VDTRead2`, `VDTTerminalPort2`, `VDTWriteBinaryWave2`, `VDTWriteBinary2`, `VDTWriteHexWave2`, `VDTWriteHex2`, `VDTWriteWave2`, `VDTWrite2`, `VDT2`, `VISAControl`, `VISARead`, `VISAReadBinary`, `VISAReadBinaryWave`, `VISAReadWave`, `VISAWrite`, `VISAWriteBinary`, `VISAWriteBinaryWave`, `VISAWriteWave`, `WaveMeanStdv`, `WaveStats`, `WaveTransform`, `wfprintf`, `WignerTransform`, `WindowFunction`, `XLLoadWave`), NameClass, nil},
			{Words(`\b`, `\b`, `abs`, `acos`, `acosh`, `AddListItem`, `AiryA`, `AiryAD`, `AiryB`, `AiryBD`, `alog`, `AnnotationInfo`, `AnnotationList`, `area`, `areaXY`, `asin`, `asinh`, `atan`, `atanh`, `atan2`, `AxisInfo`, `AxisList`, `AxisValFromPixel`, `AxonTelegraphAGetDataNum`, `AxonTelegraphAGetDataString`, `AxonTelegraphAGetDataStruct`, `AxonTelegraphGetDataNum`, `AxonTelegraphGetDataString`, `AxonTelegraphGetDataStruct`, `AxonTelegraphGetTimeoutMs`, `AxonTelegraphSetTimeoutMs`, `Base64Decode`, `Base64Encode`, `Besseli`, `Besselj`, `Besselk`, `Bessely`, `beta`, `betai`, `BinarySearch`, `BinarySearchInterp`, `binomial`, `binomialln`, `binomialNoise`, `cabs`, `CaptureHistory`, `CaptureHistoryStart`, `ceil`, `cequal`, `char2num`, `chebyshev`, `chebyshevU`, `CheckName`, `ChildWindowList`, `CleanupName`, `cmplx`, `cmpstr`, `conj`, `ContourInfo`, `ContourNameList`, `ContourNameToWaveRef`, `ContourZ`, `ControlNameList`, `ConvertTextEncoding`, `cos`, `cosh`, `cosIntegral`, `cot`, `coth`, `CountObjects`, `CountObjectsDFR`, `cpowi`, `CreationDate`, `csc`, `csch`, `CsrInfo`, `CsrWave`, `CsrWaveRef`, `CsrXWave`, `CsrXWaveRef`, `CTabList`, `DataFolderDir`, `DataFolderExists`, `DataFolderRefsEqual`, `DataFolderRefStatus`, `date`, `datetime`, `DateToJulian`, `date2secs`, `Dawson`, `defined`, `deltax`, `digamma`, `dilogarithm`, `DimDelta`, `DimOffset`, `DimSize`, `ei`, `enoise`, `equalWaves`, `erf`, `erfc`, `erfcw`, `exists`, `exp`, `expInt`, `expIntegralE1`, `expNoise`, `factorial`, `Faddeeva`, `fakedata`, `faverage`, `faverageXY`, `fDAQmx_AI_GetReader`, `fDAQmx_AO_UpdateOutputs`, `fDAQmx_ConnectTerminals`, `fDAQmx_CTR_Finished`, `fDAQmx_CTR_IsFinished`, `fDAQmx_CTR_IsPulseFinished`, `fDAQmx_CTR_ReadCounter`, `fDAQmx_CTR_ReadWithOptions`, `fDAQmx_CTR_SetPulseFrequency`, `fDAQmx_CTR_Start`, `fDAQmx_DeviceNames`, `fDAQmx_DIO_Finished`, `fDAQmx_DIO_PortWidth`, `fDAQmx_DIO_Read`, `fDAQmx_DIO_Write`, `fDAQmx_DisconnectTerminals`, `fDAQmx_ErrorString`, `fDAQmx_ExternalCalDate`, `fDAQmx_NumAnalogInputs`, `fDAQmx_NumAnalogOutputs`, `fDAQmx_NumCounters`, `fDAQmx_NumDIOPorts`, `fDAQmx_ReadChan`, `fDAQmx_ReadNamedChan`, `fDAQmx_ResetDevice`, `fDAQmx_ScanGetAvailable`, `fDAQmx_ScanGetNextIndex`, `fDAQmx_ScanStart`, `fDAQmx_ScanStop`, `fDAQmx_ScanWait`, `fDAQmx_ScanWaitWithTimeout`, `fDAQmx_SelfCalDate`, `fDAQmx_SelfCalibration`, `fDAQmx_WaveformStart`, `fDAQmx_WaveformStop`, `fDAQmx_WF_IsFinished`, `fDAQmx_WF_WaitUntilFinished`, `fDAQmx_WriteChan`, `FetchURL`, `FindDimLabel`, `FindListItem`, `floor`, `FontList`, `FontSizeHeight`, `FontSizeStringWidth`, `FresnelCos`, `FresnelSin`, `FuncRefInfo`, `FunctionInfo`, `FunctionList`, `FunctionPath`, `gamma`, `gammaEuler`, `gammaInc`, `gammaNoise`, `gammln`, `gammp`, `gammq`, `Gauss`, `Gauss1D`, `Gauss2D`, `gcd`, `GetBrowserLine`, `GetBrowserSelection`, `GetDataFolder`, `GetDataFolderDFR`, `GetDefaultFont`, `GetDefaultFontSize`, `GetDefaultFontStyle`, `GetDimLabel`, `GetEnvironmentVariable`, `GetErrMessage`, `GetFormula`, `GetIndependentModuleName`, `GetIndexedObjName`, `GetIndexedObjNameDFR`, `GetKeyState`, `GetRTErrMessage`, `GetRTError`, `GetRTLocation`, `GetRTLocInfo`, `GetRTStackInfo`, `GetScrapText`, `GetUserData`, `GetWavesDataFolder`, `GetWavesDataFolderDFR`, `GISGetAllFileFormats`, `GISSRefsAreEqual`, `GizmoInfo`, `GizmoScale`, `gnoise`, `GrepList`, `GrepString`, `GuideInfo`, `GuideNameList`, `Hash`, `hcsr`, `HDF5AttributeInfo`, `HDF5DatasetInfo`, `HDF5LibraryInfo`, `HDF5TypeInfo`, `hermite`, `hermiteGauss`, `HyperGNoise`, `HyperGPFQ`, `HyperG0F1`, `HyperG1F1`, `HyperG2F1`, `IgorInfo`, `IgorVersion`, `imag`, `ImageInfo`, `ImageNameList`, `ImageNameToWaveRef`, `IndependentModuleList`, `IndexedDir`, `IndexedFile`, `IndexToScale`, `Inf`, `Integrate1D`, `interp`, `Interp2D`, `Interp3D`, `inverseERF`, `inverseERFC`, `ItemsInList`, `JacobiCn`, `JacobiSn`, `JulianToDate`, `Laguerre`, `LaguerreA`, `LaguerreGauss`, `LambertW`, `LayoutInfo`, `leftx`, `LegendreA`, `limit`, `ListMatch`, `ListToTextWave`, `ListToWaveRefWave`, `ln`, `log`, `logNormalNoise`, `lorentzianNoise`, `LowerStr`, `MacroList`, `magsqr`, `MandelbrotPoint`, `MarcumQ`, `MatrixCondition`, `MatrixDet`, `MatrixDot`, `MatrixRank`, `MatrixTrace`, `max`, `MCC_AutoBridgeBal`, `MCC_AutoFastComp`, `MCC_AutoPipetteOffset`, `MCC_AutoSlowComp`, `MCC_AutoWholeCellComp`, `MCC_GetBridgeBalEnable`, `MCC_GetBridgeBalResist`, `MCC_GetFastCompCap`, `MCC_GetFastCompTau`, `MCC_GetHolding`, `MCC_GetHoldingEnable`, `MCC_GetMode`, `MCC_GetNeutralizationCap`, `MCC_GetNeutralizationEnable`, `MCC_GetOscKillerEnable`, `MCC_GetPipetteOffset`, `MCC_GetPrimarySignalGain`, `MCC_GetPrimarySignalHPF`, `MCC_GetPrimarySignalLPF`, `MCC_GetRsCompBandwidth`, `MCC_GetRsCompCorrection`, `MCC_GetRsCompEnable`, `MCC_GetRsCompPrediction`, `MCC_GetSecondarySignalGain`, `MCC_GetSecondarySignalLPF`, `MCC_GetSlowCompCap`, `MCC_GetSlowCompTau`, `MCC_GetSlowCompTauX20Enable`, `MCC_GetSlowCurrentInjEnable`, `MCC_GetSlowCurrentInjLevel`, `MCC_GetSlowCurrentInjSetlTime`, `MCC_GetWholeCellCompCap`, `MCC_GetWholeCellCompEnable`, `MCC_GetWholeCellCompResist`, `MCC_SelectMultiClamp700B`, `MCC_SetBridgeBalEnable`, `MCC_SetBridgeBalResist`, `MCC_SetFastCompCap`, `MCC_SetFastCompTau`, `MCC_SetHolding`, `MCC_SetHoldingEnable`, `MCC_SetMode`, `MCC_SetNeutralizationCap`, `MCC_SetNeutralizationEnable`, `MCC_SetOscKillerEnable`, `MCC_SetPipetteOffset`, `MCC_SetPrimarySignalGain`, `MCC_SetPrimarySignalHPF`, `MCC_SetPrimarySignalLPF`, `MCC_SetRsCompBandwidth`, `MCC_SetRsCompCorrection`, `MCC_SetRsCompEnable`, `MCC_SetRsCompPrediction`, `MCC_SetSecondarySignalGain`, `MCC_SetSecondarySignalLPF`, `MCC_SetSlowCompCap`, `MCC_SetSlowCompTau`, `MCC_SetSlowCompTauX20Enable`, `MCC_SetSlowCurrentInjEnable`, `MCC_SetSlowCurrentInjLevel`, `MCC_SetSlowCurrentInjSetlTime`, `MCC_SetTimeoutMs`, `MCC_SetWholeCellCompCap`, `MCC_SetWholeCellCompEnable`, `MCC_SetWholeCellCompResist`, `mean`, `median`, `min`, `mod`, `ModDate`, `MPFXEMGPeak`, `MPFXExpConvExpPeak`, `MPFXGaussPeak`, `MPFXLorenzianPeak`, `MPFXVoigtPeak`, `NameOfWave`, `NaN`, `NewFreeDataFolder`, `NewFreeWave`, `norm`, `NormalizeUnicode`, `note`, `NumberByKey`, `numpnts`, `numtype`, `NumVarOrDefault`, `num2char`, `num2istr`, `num2str`, `NVAR_Exists`, `OperationList`, `PadString`, `PanelResolution`, `ParamIsDefault`, `ParseFilePath`, `PathList`, `pcsr`, `Pi`, `PICTInfo`, `PICTList`, `PixelFromAxisVal`, `pnt2x`, `poissonNoise`, `poly`, `PolygonArea`, `poly2D`, `PossiblyQuoteName`, `ProcedureText`, `p2rect`, `qcsr`, `real`, `RemoveByKey`, `RemoveEnding`, `RemoveFromList`, `RemoveListItem`, `ReplaceNumberByKey`, `ReplaceString`, `ReplaceStringByKey`, `rightx`, `round`, `r2polar`, `sawtooth`, `scaleToIndex`, `ScreenResolution`, `sec`, `sech`, `Secs2Date`, `Secs2Time`, `SelectNumber`, `SelectString`, `SetEnvironmentVariable`, `sign`, `sin`, `sinc`, `sinh`, `sinIntegral`, `SortList`, `SpecialCharacterInfo`, `SpecialCharacterList`, `SpecialDirPath`, `SphericalBessJ`, `SphericalBessJD`, `SphericalBessY`, `SphericalBessYD`, `SphericalHarmonics`, `SQLAllocHandle`, `SQLAllocStmt`, `SQLBinaryWavesToTextWave`, `SQLBindCol`, `SQLBindParameter`, `SQLBrowseConnect`, `SQLBulkOperations`, `SQLCancel`, `SQLCloseCursor`, `SQLColAttributeNum`, `SQLColAttributeStr`, `SQLColumnPrivileges`, `SQLColumns`, `SQLConnect`, `SQLDataSources`, `SQLDescribeCol`, `SQLDescribeParam`, `SQLDisconnect`, `SQLDriverConnect`, `SQLDrivers`, `SQLEndTran`, `SQLError`, `SQLExecDirect`, `SQLExecute`, `SQLFetch`, `SQLFetchScroll`, `SQLForeignKeys`, `SQLFreeConnect`, `SQLFreeEnv`, `SQLFreeHandle`, `SQLFreeStmt`, `SQLGetConnectAttrNum`, `SQLGetConnectAttrStr`, `SQLGetCursorName`, `SQLGetDataNum`, `SQLGetDataStr`, `SQLGetDescFieldNum`, `SQLGetDescFieldStr`, `SQLGetDescRec`, `SQLGetDiagFieldNum`, `SQLGetDiagFieldStr`, `SQLGetDiagRec`, `SQLGetEnvAttrNum`, `SQLGetEnvAttrStr`, `SQLGetFunctions`, `SQLGetInfoNum`, `SQLGetInfoStr`, `SQLGetStmtAttrNum`, `SQLGetStmtAttrStr`, `SQLGetTypeInfo`, `SQLMoreResults`, `SQLNativeSql`, `SQLNumParams`, `SQLNumResultCols`, `SQLNumResultRowsIfKnown`, `SQLNumRowsFetched`, `SQLParamData`, `SQLPrepare`, `SQLPrimaryKeys`, `SQLProcedureColumns`, `SQLProcedures`, `SQLPutData`, `SQLReinitialize`, `SQLRowCount`, `SQLSetConnectAttrNum`, `SQLSetConnectAttrStr`, `SQLSetCursorName`, `SQLSetDescFieldNum`, `SQLSetDescFieldStr`, `SQLSetDescRec`, `SQLSetEnvAttrNum`, `SQLSetEnvAttrStr`, `SQLSetPos`, `SQLSetStmtAttrNum`, `SQLSetStmtAttrStr`, `SQLSpecialColumns`, `SQLStatistics`, `SQLTablePrivileges`, `SQLTables`, `SQLTextWaveToBinaryWaves`, `SQLTextWaveTo2DBinaryWave`, `SQLUpdateBoundValues`, `SQLXOPCheckState`, `SQL2DBinaryWaveToTextWave`, `sqrt`, `StartMSTimer`, `StatsBetaCDF`, `StatsBetaPDF`, `StatsBinomialCDF`, `StatsBinomialPDF`, `StatsCauchyCDF`, `StatsCauchyPDF`, `StatsChiCDF`, `StatsChiPDF`, `StatsCMSSDCDF`, `StatsCorrelation`, `StatsDExpCDF`, `StatsDExpPDF`, `StatsErlangCDF`, `StatsErlangPDF`, `StatsErrorPDF`, `StatsEValueCDF`, `StatsEValuePDF`, `StatsExpCDF`, `StatsExpPDF`, `StatsFCDF`, `StatsFPDF`, `StatsFriedmanCDF`, `StatsGammaCDF`, `StatsGammaPDF`, `StatsGeometricCDF`, `StatsGeometricPDF`, `StatsGEVCDF`, `StatsGEVPDF`, `StatsHyperGCDF`, `StatsHyperGPDF`, `StatsInvBetaCDF`, `StatsInvBinomialCDF`, `StatsInvCauchyCDF`, `StatsInvChiCDF`, `StatsInvCMSSDCDF`, `StatsInvDExpCDF`, `StatsInvEValueCDF`, `StatsInvExpCDF`, `StatsInvFCDF`, `StatsInvFriedmanCDF`, `StatsInvGammaCDF`, `StatsInvGeometricCDF`, `StatsInvKuiperCDF`, `StatsInvLogisticCDF`, `StatsInvLogNormalCDF`, `StatsInvMaxwellCDF`, `StatsInvMooreCDF`, `StatsInvNBinomialCDF`, `StatsInvNCChiCDF`, `StatsInvNCFCDF`, `StatsInvNormalCDF`, `StatsInvParetoCDF`, `StatsInvPoissonCDF`, `StatsInvPowerCDF`, `StatsInvQCDF`, `StatsInvQpCDF`, `StatsInvRayleighCDF`, `StatsInvRectangularCDF`, `StatsInvSpearmanCDF`, `StatsInvStudentCDF`, `StatsInvTopDownCDF`, `StatsInvTriangularCDF`, `StatsInvUsquaredCDF`, `StatsInvVonMisesCDF`, `StatsInvWeibullCDF`, `StatsKuiperCDF`, `StatsLogisticCDF`, `StatsLogisticPDF`, `StatsLogNormalCDF`, `StatsLogNormalPDF`, `StatsMaxwellCDF`, `StatsMaxwellPDF`, `StatsMedian`, `StatsMooreCDF`, `StatsNBinomialCDF`, `StatsNBinomialPDF`, `StatsNCChiCDF`, `StatsNCChiPDF`, `StatsNCFCDF`, `StatsNCFPDF`, `StatsNCTCDF`, `StatsNCTPDF`, `StatsNormalCDF`, `StatsNormalPDF`, `StatsParetoCDF`, `StatsParetoPDF`, `StatsPermute`, `StatsPoissonCDF`, `StatsPoissonPDF`, `StatsPowerCDF`, `StatsPowerNoise`, `StatsPowerPDF`, `StatsQCDF`, `StatsQpCDF`, `StatsRayleighCDF`, `StatsRayleighPDF`, `StatsRectangularCDF`, `StatsRectangularPDF`, `StatsRunsCDF`, `StatsSpearmanRhoCDF`, `StatsStudentCDF`, `StatsStudentPDF`, `StatsTopDownCDF`, `StatsTriangularCDF`, `StatsTriangularPDF`, `StatsTrimmedMean`, `StatsUSquaredCDF`, `StatsVonMisesCDF`, `StatsVonMisesNoise`, `StatsVonMisesPDF`, `StatsWaldCDF`, `StatsWaldPDF`, `StatsWeibullCDF`, `StatsWeibullPDF`, `StopMSTimer`, `StringByKey`, `stringCRC`, `StringFromList`, `StringList`, `stringmatch`, `strlen`, `strsearch`, `StrVarOrDefault`, `str2num`, `StudentA`, `StudentT`, `sum`, `SVAR_Exists`, `TableInfo`, `TagVal`, `TagWaveRef`, `tan`, `tango_close_device`, `tango_command_inout`, `tango_compute_image_proj`, `tango_get_dev_attr_list`, `tango_get_dev_black_box`, `tango_get_dev_cmd_list`, `tango_get_dev_status`, `tango_get_dev_timeout`, `tango_get_error_stack`, `tango_open_device`, `tango_ping_device`, `tango_read_attribute`, `tango_read_attributes`, `tango_reload_dev_interface`, `tango_resume_attr_monitor`, `tango_set_attr_monitor_period`, `tango_set_dev_timeout`, `tango_start_attr_monitor`, `tango_stop_attr_monitor`, `tango_suspend_attr_monitor`, `tango_write_attribute`, `tango_write_attributes`, `tanh`, `TDMAddChannel`, `TDMAddGroup`, `TDMAppendDataValues`, `TDMAppendDataValuesTime`, `TDMChannelPropertyExists`, `TDMCloseChannel`, `TDMCloseFile`, `TDMCloseGroup`, `TDMCreateChannelProperty`, `TDMCreateFile`, `TDMCreateFileProperty`, `TDMCreateGroupProperty`, `TDMFilePropertyExists`, `TDMGetChannelPropertyNames`, `TDMGetChannelPropertyNum`, `TDMGetChannelPropertyStr`, `TDMGetChannelPropertyTime`, `TDMGetChannelPropertyType`, `TDMGetChannels`, `TDMGetChannelStringPropertyLen`, `TDMGetDataType`, `TDMGetDataValues`, `TDMGetDataValuesTime`, `TDMGetFilePropertyNames`, `TDMGetFilePropertyNum`, `TDMGetFilePropertyStr`, `TDMGetFilePropertyTime`, `TDMGetFilePropertyType`, `TDMGetFileStringPropertyLen`, `TDMGetGroupPropertyNames`, `TDMGetGroupPropertyNum`, `TDMGetGroupPropertyStr`, `TDMGetGroupPropertyTime`, `TDMGetGroupPropertyType`, `TDMGetGroups`, `TDMGetGroupStringPropertyLen`, `TDMGetLibraryErrorDescription`, `TDMGetNumChannelProperties`, `TDMGetNumChannels`, `TDMGetNumDataValues`, `TDMGetNumFileProperties`, `TDMGetNumGroupProperties`, `TDMGetNumGroups`, `TDMGroupPropertyExists`, `TDMOpenFile`, `TDMOpenFileEx`, `TDMRemoveChannel`, `TDMRemoveGroup`, `TDMReplaceDataValues`, `TDMReplaceDataValuesTime`, `TDMSaveFile`, `TDMSetChannelPropertyNum`, `TDMSetChannelPropertyStr`, `TDMSetChannelPropertyTime`, `TDMSetDataValues`, `TDMSetDataValuesTime`, `TDMSetFilePropertyNum`, `TDMSetFilePropertyStr`, `TDMSetFilePropertyTime`, `TDMSetGroupPropertyNum`, `TDMSetGroupPropertyStr`, `TDMSetGroupPropertyTime`, `TextEncodingCode`, `TextEncodingName`, `TextFile`, `ThreadGroupCreate`, `ThreadGroupGetDF`, `ThreadGroupGetDFR`, `ThreadGroupRelease`, `ThreadGroupWait`, `ThreadProcessorCount`, `ThreadReturnValue`, `ticks`, `time`, `TraceFromPixel`, `TraceInfo`, `TraceNameList`, `TraceNameToWaveRef`, `TrimString`, `trunc`, `UniqueName`, `UnPadString`, `UnsetEnvironmentVariable`, `UpperStr`, `URLDecode`, `URLEncode`, `VariableList`, `Variance`, `vcsr`, `viAssertIntrSignal`, `viAssertTrigger`, `viAssertUtilSignal`, `viClear`, `viClose`, `viDisableEvent`, `viDiscardEvents`, `viEnableEvent`, `viFindNext`, `viFindRsrc`, `viGetAttribute`, `viGetAttributeString`, `viGpibCommand`, `viGpibControlATN`, `viGpibControlREN`, `viGpibPassControl`, `viGpibSendIFC`, `viIn8`, `viIn16`, `viIn32`, `viLock`, `viMapAddress`, `viMapTrigger`, `viMemAlloc`, `viMemFree`, `viMoveIn8`, `viMoveIn16`, `viMoveIn32`, `viMoveOut8`, `viMoveOut16`, `viMoveOut32`, `viOpen`, `viOpenDefaultRM`, `viOut8`, `viOut16`, `viOut32`, `viPeek8`, `viPeek16`, `viPeek32`, `viPoke8`, `viPoke16`, `viPoke32`, `viRead`, `viReadSTB`, `viSetAttribute`, `viSetAttributeString`, `viStatusDesc`, `viTerminate`, `viUnlock`, `viUnmapAddress`, `viUnmapTrigger`, `viUsbControlIn`, `viUsbControlOut`, `viVxiCommandQuery`, `viWaitOnEvent`, `viWrite`, `VoigtFunc`, `VoigtPeak`, `WaveCRC`, `WaveDims`, `WaveExists`, `WaveHash`, `WaveInfo`, `WaveList`, `WaveMax`, `WaveMin`, `WaveName`, `WaveRefIndexed`, `WaveRefIndexedDFR`, `WaveRefsEqual`, `WaveRefWaveToList`, `WaveTextEncoding`, `WaveType`, `WaveUnits`, `WhichListItem`, `WinList`, `WinName`, `WinRecreation`, `WinType`, `wnoise`, `xcsr`, `XWaveName`, `XWaveRefFromTrace`, `x2pnt`, `zcsr`, `ZernikeR`, `zeromq_client_connect`, `zeromq_client_connect`, `zeromq_client_recv`, `zeromq_client_recv`, `zeromq_client_send`, `zeromq_client_send`, `zeromq_handler_start`, `zeromq_handler_start`, `zeromq_handler_stop`, `zeromq_handler_stop`, `zeromq_server_bind`, `zeromq_server_bind`, `zeromq_server_recv`, `zeromq_server_recv`, `zeromq_server_send`, `zeromq_server_send`, `zeromq_set`, `zeromq_set`, `zeromq_stop`, `zeromq_stop`, `zeromq_test_callfunction`, `zeromq_test_callfunction`, `zeromq_test_serializeWave`, `zeromq_test_serializeWave`, `zeta`), NameFunction, nil},
			{`^#(include|pragma|define|undef|ifdef|ifndef|if|elif|else|endif)`, NameDecorator, nil},
			{`[^a-z"/]+$`, Text, nil},
			{`.`, Text, nil},
			{`\n|\r`, Text, nil},
		},
	},
))
