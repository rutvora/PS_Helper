# PS Helper  


_Note 1:_ This script needs modification for every PS. Modification is detailed below:  
In line 297 of the script, BatchIdFor has to be changed to current batch. For PS-I, batch 10 is PS-I in 2020.  
In line 297 of the script, PSTypeFor has to be changed to PS-I or PS-II based on the requirement.  
That's all that is required for it to work well  

**_Instructions:_**
1. Open Google Chrome
2. Press F12 to show debugging mode
3. Open PSD site and Login (keep the debug thingy open)
4. After Login, Go to "Network" tab in the debug thingy
5. Select Login.aspx from the list of files
6. Go to "Headers" and find "Cookies" written in bold
7. Copy the cookie's value
8. Go to the folder where you downloaded/compiled the app (using terminal)
9. See Commands below

**_Commands:_**  
_Note:_ paste whatever you copied from chrome where I've written cookie  
_Windows Users: _ Use scraper.exe instead of ./scraper  
_Mac Users: _ Use ./scraper_mac instead of ./scraper  

Create csv file  
`./scraper -g "cookie"` (with quotes)  
Update List (from the same csv file)  
`./scraper -u "cookie"` (with quotes)

