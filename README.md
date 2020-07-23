# PS Helper  


_Note 1:_ This script needs modification for every PS. Modification is detailed below:  
In line 217 of the script, BatchIdFor has to be changed to current batch. For PS-I, batch 10 is PS-I in 2020.  
In line 217 of the script, PSTypeFor has to be changed to PS-I or PS-II based on the requirement.  
That's all that is required for it to work well  

**_Instructions:_**
1. (To be run only the first time you download the script) chmod +x ./scraper *
2. Open Google Chrome
3. Press F12 to show debugging mode
4. Open PSD site and Login (keep the debug thingy open)
5. After Login, Go to "Network" tab in the debug thingy
6. Select Login.aspx from the list of files
7. Go to "Headers" and find "Cookies" written in bold
8. Copy the cookie's value
9. Go to the folder where you downloaded/compiled the app (using terminal)
10. See Commands below

\* Not for Windows users

**_Commands:_**  
_Note:_ paste whatever you copied from chrome where I've written cookie  
_Windows Users: _ Use scraper.exe instead of ./scraper  
_Mac Users: _ Use ./scraper_mac instead of ./scraper  

Create csv file  
`./scraper "cookie"` (with quotes)  