export const provideMock = <T>(value: unknown) => value as T;
export const tick = () =>
    new Promise((resolve) => {
        setTimeout(resolve, 0);
    });
