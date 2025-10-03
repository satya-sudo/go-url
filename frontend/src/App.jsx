import { BrowserRouter, Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import ProtectedRoute from "./components/ProtectedRoute";
import Signup from "./pages/Signup";
import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";
import Stats from "./pages/Stats";
import "./styles.css";

function App() {
    return (
        <div className="app-container">
            <BrowserRouter>
                <Navbar />
                <main className="main-content">
                    <Routes>
                        <Route path="/signup" element={<Signup />} />
                        <Route path="/login" element={<Login />} />
                        <Route
                            path="/"
                            element={
                                <ProtectedRoute>
                                    <Dashboard />
                                </ProtectedRoute>
                            }
                        />
                        <Route
                            path="/stats"
                            element={
                                <ProtectedRoute>
                                    <Stats />
                                </ProtectedRoute>
                            }
                        />
                    </Routes>
                </main>
            </BrowserRouter>
        </div>
    );
}

export default App;
