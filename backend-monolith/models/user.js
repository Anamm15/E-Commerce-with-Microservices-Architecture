import db from "../config/database.js";
import { DataTypes } from "sequelize";

const Users = db.define('user', {
    name: {
        type: DataTypes.STRING,
        allowNull: false
    },
    bio: {
        type: DataTypes.TEXT,
    },
    gender: {
        type: DataTypes.CHAR,
        allowNull: false
    },
    birth_place: {
        type: DataTypes.STRING,
        allowNull: false
    },
    birth_date: {
        type: DataTypes.DATE,
        allowNull: false
    },
    address: {
        type: DataTypes.STRING,
    },
    phone_num: {
        type: DataTypes.STRING(20),
    },
    email: {
        type: DataTypes.STRING,
        allowNull: false,
        unique: true
    },
    username: {
        type: DataTypes.STRING,
        allowNull: false,
        unique: true
    },
    password: {
        type: DataTypes.STRING,
        allowNull: false
    },
    refresh_token: {
        type: DataTypes.STRING
    },
    photo: {
        type: DataTypes.TEXT
    }
}, {
    freezeTableName: true,
    timestamps: true,
    paranoid: true
});

export default Users;