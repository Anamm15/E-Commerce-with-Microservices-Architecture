import { Sequelize } from "sequelize";

const db = new Sequelize('anamstore', 'root', '', {
    dialect: 'mysql',
    host: 'localhost'
});

export default db;