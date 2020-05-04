// CLIENT
export const APP_ROUTES = {
    root: '/',
    about: '/about/',
    userPage: '/user/:id/home/:page?'
};

// SERVER
export const APP_BACKEND_BASE = 'http://localhost:8081/api/kadvisor/:uid';
export const APP_CLASS_ENDPOINT = '/class';
export const APP_ENTRY_ENDPOINT = '/entry';
export const APP_LOGIN_ENDPOINT = {
    login: '/login',
    logout: '/logout'
};
