import db from "../config/database.js";
import { DataTypes } from "sequelize";

const Stores = db.define('store', {
    name: {
        type: DataTypes.STRING,
        allowNull: false
    },
    description: {
        type: DataTypes.TEXT,
    },
    domicile: {
        type: DataTypes.STRING
    },
    address: {
        type: DataTypes.STRING
    },
    phone_num: {
        type: DataTypes.STRING
    },
    email: {
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

export default Stores;