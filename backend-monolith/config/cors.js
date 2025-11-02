const corsOptions = {
    origin: 'http://localhost:5172',
    credentials: true,
    optionSuccessStatus: 204,
    methods: 'GET,HEAD,PUT,PATCH,POST,DELETE'
}

export default corsOptions;