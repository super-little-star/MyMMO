@SET No_Sheet_EXCEL_FOLDER=.\excel\NoSheet
@SET EXCEL_FOLDER=.\excel\Sheet
@SET JSON_FOLDER=.\json
@SET EXE=.\excel2json\excel2json.exe

@ECHO Converting excel files in folder %EXCEL_FOLDER% ...
for /f "delims=" %%i in ('dir /b /a-d /s %EXCEL_FOLDER%\*.xlsx') do (
    @echo   processing %%~nxi 
    @CALL %EXE% --excel %EXCEL_FOLDER%\%%~nxi --json %JSON_FOLDER%\%%~ni.json --header 3 --s
)

@ECHO Converting excel files in folder %No_Sheet_EXCEL_FOLDER% ...
for /f "delims=" %%i in ('dir /b /a-d /s %No_Sheet_EXCEL_FOLDER%\*.xlsx') do (
    @echo   processing %%~nxi 
    @CALL %EXE% --excel %No_Sheet_EXCEL_FOLDER%\%%~nxi --json %JSON_FOLDER%\%%~ni.json --header 3 
)


xcopy .\json ..\Client\ClientProject\Assets\Resources\Data
xcopy .\json ..\Server\Src\data
@pause