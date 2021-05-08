module.exports = {
    displayName: 'kadvisor-app',
    preset: '../../jest.preset.js',
    transform: {
        '^(?!.*\\.(js|jsx|ts|tsx|css|json)$)': '@nrwl/react/plugins/jest',
        '^.+\\.[tj]sx?$': 'babel-jest'
    },
    moduleFileExtensions: ['ts', 'tsx', 'js', 'jsx'],
    coverageDirectory: '../../coverage/apps/kadvisor-app',
    setupFiles: ['./../../node_modules/jest-canvas-mock/lib/index.js'],
    setupFilesAfterEnv: ['<rootDir>/enzyme.setup.js']
};
