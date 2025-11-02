import { getCommentPerProduct as getCommentPerProductFromService, addComment as addCommentFromService } from "../services/commentService.js";
import { convertImagePathToLink } from "../utils/utils.js";

export const getCommentPerProduct = async(req, res) => {
    try {
        const comments = await getCommentPerProductFromService(req.params.id);
        if (!comments) return res.sendStatus(404);
        comments.map((comment) => comment.user.photo = convertImagePathToLink(req.protocol, req.get('host'), comment.user.photo));
        res.json(comments);
    } catch (error) {
        res.status(500).json({message: "Internal server error"});
    }
}

export const addComment = async(req, res) => {
    try {
        await addCommentFromService(req.body);
        res.json({message: "Komen berhasil ditambahkan"});
    } catch (error) {
        res.status(500).json({message: "Internal server error"});
    }
}