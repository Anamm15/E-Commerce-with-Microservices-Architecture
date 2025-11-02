import jwt from "jsonwebtoken";
import { getUserByRefreshToken, getUserByUsername, updateRefreshToken } from "../services/userService.js";
import bcrypt from 'bcrypt';

export const newToken = async(req,res) => {
    try {
        const refreshToken = req.cookies.token;
        if(!refreshToken) return res.sendStatus(401);
        const user = await getUserByRefreshToken(refreshToken);
        if(!user) return res.sendStatus(403);
        jwt.verify(refreshToken, process.env.REFRESH_KEY, (err, decoded) => {
            if(err) return res.sendStatus(403);
            const newToken = jwt.sign({
                id: user.id,
            }, process.env.SECRET_KEY, {
                expiresIn: '20m'
            });
            res.json({accessToken: newToken});
        })
    } catch (error) {
        res.status(500).json({ message: "Internal server error"});
    }
}

export const login = async(req, res) => {
    try {
        const { username, password } = req.body;

        const user = await getUserByUsername(username);
        
        if(!user) return res.status(404).json({ message: 'User atau Password salah' });
        
        const isMatch = await bcrypt.compare(password, user.password);
        
        if(!isMatch) return res.status(401).json({ message: 'User atau Password salah' });
        
        const token = jwt.sign({ id: user.id }, process.env.SECRET_KEY, {
            expiresIn: '20m'
        });

        const refreshToken = jwt.sign({ id: user.id }, process.env.REFRESH_KEY, {
            expiresIn: '1d'
        });

        updateRefreshToken(user.id, refreshToken);

        res.cookie('token', refreshToken, {
            httpOnly: true,
            maxAge: 60 * 60 * 24 * 1000,
        });

        res.cookie('session', 'sdadf', {
            maxAge: 60 * 60 * 24 * 1000   
        })

        res.status(200).json({accessToken : token});
    } catch (error) {
        res.status(500).json({ message: 'Internal server error' });
    }
}

export const logout = async(req, res) => {
    try {
        const refreshToken = req.cookies.token;
        if(!refreshToken) return res.sendStatus(204);
        // const user = await Users.findOne({ where: { refresh_token: refreshToken } });
        // if(!user) return res.sendStatus(204);
        // await Users.update({
        //     refresh_token: null
        // }, {
        //     where: {refresh_token: refreshToken}
        // });
        res.clearCookie('token');
        res.clearCookie('session');
        return res.sendStatus(200);
    } catch (error) {
        res.status(500).json({message: "Internal server error"});
    }
}