# PS Helper  
_Note 1:_ There is a separate branch to fetch project details. Please check it out.  
_Note 2:_ This script seems to be working for both PS-I and PS-II. If you face any trouble, ping me on facebook.  
_Note 3:_ You can find the binaries in the [releases](https://github.com/rutvora/PS_Helper/releases). Make sure you are downloading binaries of the master branch.  

**_Video Tutorial Link:_** https://youtu.be/I7PGqZqXOzA  

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
_Note:_ paste whatever you copied from chrome where I've written cookie and make sure you are using the correct binary  

Create csv file  
`./<binary> -g "cookie"` (with quotes, don't forget to replace <binary> with the binary that you downloaded)  
Update List (from the same csv file)  
`./<binary> -u "cookie"` (with quotes ,don't forget to replace <binary> with the binary that you downloaded)
 
