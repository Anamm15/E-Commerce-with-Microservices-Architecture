import { getCategories as getCategoriesFromService } from "../services/categoryService.js";


export const getCategories = async (req, res) => {
    try {
        const categories = await getCategoriesFromService();

        if (!categories || categories.length === 0) {
            return res.sendStatus(404);
        }
        
        res.json(categories);
    } catch (error) {
        res.status(500).json({message: "Internal server error"});
    }
}