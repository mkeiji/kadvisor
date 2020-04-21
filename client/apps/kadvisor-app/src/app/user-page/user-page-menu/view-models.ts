export const UserPageMenuListTypes = {
    primary: 'primary',
    secondary: 'secondary',
    divider: 'divider'
};

export interface UserPageMenuListObject {
    itemText?: string;
    pagePath?: string;
    type: string;
    iconComponent: JSX.Element;
}
