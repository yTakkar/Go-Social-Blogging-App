# Go-Social-Blogging-App

A Social-Blogging App developed with Golang & Mysql!!

### Updated Version
Click [here](https://github.com/yTakkar/Go-Mini-Social-Network) for mini-social-network created with Golang!!

# Quick Links
1. [Screenshots](#screenshots)
2. [Requirements](#requirements)
3. [Usage](#usage)

# Screenshots
![alt text](https://raw.githubusercontent.com/yTakkar/Go-Social-Blogging-App/master/screenshots/Snap%202017-09-05%20at%2015.40.35.png)
![alt text](https://raw.githubusercontent.com/yTakkar/Go-Social-Blogging-App/master/screenshots/Snap%202017-09-05%20at%2015.40.46.png)
![alt text](https://raw.githubusercontent.com/yTakkar/Go-Social-Blogging-App/master/screenshots/Snap%202017-09-05%20at%2015.40.54.png)
![alt text](https://raw.githubusercontent.com/yTakkar/Go-Social-Blogging-App/master/screenshots/Snap%202017-09-05%20at%2015.41.31.png)
![alt text](https://raw.githubusercontent.com/yTakkar/Go-Social-Blogging-App/master/screenshots/Snap%202017-09-05%20at%2015.41.39.png)
![alt text](https://raw.githubusercontent.com/yTakkar/Go-Social-Blogging-App/master/screenshots/Snap%202017-09-05%20at%2015.41.10.png)
![alt text](https://raw.githubusercontent.com/yTakkar/Go-Social-Blogging-App/master/screenshots/Snap%202017-09-05%20at%2015.41.47.png)

# Requirements
1. Make sure you keep this project inside `src/` of your Golang project folder.
2. Following packages should be installed.
    1. [httprouter](https://github.com/julienschmidt/httprouter)
    2. [negroni](https://github.com/urfave/negroni)
    3. [checkmail](https://github.com/badoux/checkmail)
    4. [MySQL driver](https://github.com/go-sql-driver/mysql)
    5. [bcrypt](https://golang.org/x/crypto/bcrypt)
    6. [sessions](https://github.com/gorilla/sessions)
    7. [godotenv](https://github.com/joho/godotenv)

# Usage

1. Open PHPMyAdmin, create a db & import `db.sql`.

2. Install all the dependencies with npm or Yarn.
```javascript
npm install
```
or
```javascript
yarn
```

3. Create a `.env` file & insert the above code. Replace values with yours!!
```javascript
PORT=YOUR PORT (default: 2280) [STRING]
SESSION_SECRET=ANYTHING SECRET [STRING]
DB_USER=DB_USER [STRING]
DB_PASSWORD=DB PASSWORD [STRING]
DB=DB YOU JUST CREATE [STRING]
```

4. My root folder name is `Go-Social-Blog-App`, if yours is different then go ahead & change it as it used for imports. It can be done easily by searching the whole project.

5. Now run the app.
```javascript
npm run start
```

6. Run the app in browser.
```javascript
localhost:[PORT] PORT=2280 (By default)
```
