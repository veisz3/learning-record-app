import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import { AppBar, Toolbar, Typography, Container, Button, CssBaseline, ThemeProvider, createTheme } from '@mui/material';
import LearningForm from './components/LearningForm';
import LearningList from './components/LearningList';

const theme = createTheme({
    palette: {
        mode: 'light',
        primary: {
            main: '#2196f3',
        },
        background: {
            default: '#f5f5f5',
        },
    },
});

const App: React.FC = () => {
    return (
        <ThemeProvider theme={theme}>
            <CssBaseline />
            <Router>
                <AppBar position="static">
                    <Toolbar>
                        <Typography variant="h6" style={{ flexGrow: 1 }}>
                            Learning Tracker
                        </Typography>
                        <Button color="inherit" component={Link} to="/">
                            Home
                        </Button>
                        <Button color="inherit" component={Link} to="/list">
                            List
                        </Button>
                    </Toolbar>
                </AppBar>
                <Container sx={{ mt: 4 }}>
                    <Routes>
                        <Route path="/" element={<LearningForm />} />
                        <Route path="/list" element={<LearningList />} />
                    </Routes>
                </Container>
            </Router>
        </ThemeProvider>
    );
}

export default App
