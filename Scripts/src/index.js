import React from 'react';
import ReactDOM from 'react-dom';
import Header from "./components/Header"
import Body from "./components/Body"
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';

const theme = createMuiTheme({
    palette: {
        type: 'dark',
        primary: {
            light: '#484848',
            main: '#212121',
            dark: '#000000',
            contrastText: '#ffffff',
        },
        secondary: {
            light: '#819ca9',
            main: '#546e7a',
            dark: '#29434e',
            contrastText: '#ffffff',
        },
    },
});

ReactDOM.render(
    <ThemeProvider theme={theme}>
        <Header />
        <Body />
    </ThemeProvider>
    ,
    document.getElementById('root')
)


