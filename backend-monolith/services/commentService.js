import { Comments, Users } from "../models/index.js";
import { convertImagePathToLink } from "../utils/utils.js";


export const getCommentPerProduct = async(id) => {
    return await Comments.findAll({
        where: { 
            productId: id
        },
        include: {
            model: Users,
            attributes: ['username', 'photo']
        }
    });
}

export const addComment = async(data) => {
    return await Comments.create(data);
}