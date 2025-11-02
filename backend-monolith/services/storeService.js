import { Stores } from "../models/index.js";


export const getStores = async() => {
    return await Stores.findAll();
}

export const getStoreById = async(id) => {
    return await Stores.findByPk(id);
}

export const getStoreByUserId = async(userId) => {
    return await Stores.findOne({ where: { userId: parseInt(userId) }});
}

export const addStore = async(data) => {
    return await Stores.create({
        name: data.name, 
        description: data.description,
        domicile: data.domicile,
        address: data.address,
        phone_num: data.phone_num,
        email: data.email,
        photo: data.photo,
        userId: parseInt(data.userId)
    });
}