// CLIENT
export const APP_ROUTES = {
    root: '/',
    about: '/about/',
    userPage: '/user/:id/home/:page?'
};

export const APP_PAGES = {
    home: 'home',
    dashboard: 'dashboard',
    entries: 'entries',
    reports: 'reports',
    settings: 'settings'
};

// SERVER
export const APP_DOMAIN = `${process.env.NX_APP_DOMAIN}`;
export const APP_BACKEND = `${process.env.NX_APP_DOMAIN}/api`;
export const APP_BACKEND_BASE = `${APP_BACKEND}/kadvisor/:uid`;
export const APP_LOOKUP_ENDPOINT = '/lookup';
export const APP_CLASS_ENDPOINT = '/class';
export const APP_ENTRY_ENDPOINT = '/entry';
export const APP_REPORT_ENDPOINT = '/report';
export const APP_REPORT_AVAILABLE_ENDPOINT = '/reportavailable';
export const APP_FORECAST_ENDPOINT = '/forecast';
export const APP_FORECAST_ENTRY_ENDPOINT = '/forecastentry';
export const APP_LOGIN_ENDPOINT = {
    auth: '/auth',
    login: '/login',
    logout: '/logout'
};
