import { Categories } from "../models/index.js"

export const getCategories = async() => {
    return await Categories.findAll();
}