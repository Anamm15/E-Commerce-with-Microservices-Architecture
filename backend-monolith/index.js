import express from 'express';
import router from './routes/router.js';
import db from './config/database.js';
import cors from 'cors';
import corsOptions from './config/cors.js';
import dotenv from 'dotenv';
import cookieParser from 'cookie-parser';
import bodyParser from 'body-parser';
// import path from 'path';

dotenv.config();
db.authenticate()
.then(() => {
  console.log('Database connection has been established successfully.');
})
.catch(err => {
  console.error('Unable to connect to the database:', err);
});
  
const app = express();
app.use('/public', express.static('public'));
app.use(cors(corsOptions));
app.use(cookieParser());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.json());
app.use(router);

app.listen(process.env.PORT, () => {
  console.log(`Server started on port ${process.env.PORT}`);
});