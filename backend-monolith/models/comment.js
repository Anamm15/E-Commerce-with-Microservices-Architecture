import db from "../config/database.js";
import { DataTypes } from "sequelize";

const Comments = db.define('comment', {
    comment: {
        type: DataTypes.TEXT,
        allowNull: false
    },
    photo: {
        type: DataTypes.BLOB('long')
    }
}, {
    freezeTableName: true,
    timestamps: true,
    paranoid: true
});

export default Comments;