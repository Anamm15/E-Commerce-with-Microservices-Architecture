import { Products } from "../models/index.js"

export const getProducts = async(id) => {
    const products = await Products.findAll({
        attributes: {
            exclude: ['createdAt', 'updatedAt', 'deletedAt']
        }
    });
    return products;
}

export const getProduct = async(id, t) => {
    const product = await Products.findOne({ where: { id }, transaction: t });
    return product;
}

export const addProduct = async(data) => {
    return await Products.create({
        name: data.name, 
        description: data.description,
        price: parseFloat(data.price),
        stock: parseInt(data.stock),
        categoryId: parseInt(data.categoryId), 
        storeId: parseInt(data.storeId),
        photo: data.photo,
        rating: 0,
        sold: 0
    });
}

export const updateStockProduct = async(product, qty, t) => {
    return await Products.update({stok: (produk.stock - qty),
        terjual: (product.terjual + qty)}, {where: {id: product.id}, transaction: t});   
}
