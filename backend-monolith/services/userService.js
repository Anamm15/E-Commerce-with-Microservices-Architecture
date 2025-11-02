import { Users } from "../models/index.js";
import bcrypt from "bcrypt";
import jwt from "jsonwebtoken"
import fs from "fs";

export const getUsers = async() => {
    const users = await Users.findAll();
    return users
}

export const getUser = async(id) => {
    const user = await Users.findByPk(id);
    return user
}

export const register = async(data) => {
    const photo = 'public/img/defaultPhotoProfile.jpg';
    const salt = await bcrypt.genSalt(10);
    const hashedPassword = await bcrypt.hash(data.password, salt);

    return await Users.create({ 
        name: data.name,
        username: data.username,
        email: data.email,
        password: hashedPassword,
        gender: data.gender,
        birth_place: data.birth_place,
        birth_date: data.birth_date,
        address: data.address,
        phone_num: data.phone_num,
        photo: photo
    });
}

export const getUserByUsername = async(username) => {
    const user = await Users.findOne({ where: {username} });
    return user;
}

export const getUserByRefreshToken =  async(refreshToken) => {
    const user = await Users.findOne({ where: { refresh_token: refreshToken }});
    return user;
}

export const updateRefreshToken = async(id, refreshToken) => {
    return Users.update({
        refresh_token: refreshToken
    }, {
        where: { id }
    })
}

export const updatePhotoProfile = async(data) => {

}

export const deleteUser = async(id) => {
    return Users.destroy({where: {id}});
}