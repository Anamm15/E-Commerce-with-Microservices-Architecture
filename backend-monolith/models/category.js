import db from "../config/database.js";
import { DataTypes } from "sequelize";

const Categories = db.define('category', {
    category: {
        type: DataTypes.STRING,
    },
    sub_category: {
        type: DataTypes.STRING,
        allowNull: true
    }
}, {
    freezeTableName: true,
    timestamps: true,
    paranoid: true
});

export default Categories;