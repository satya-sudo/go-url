import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api/axios";

export default function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();
        setError("");
        try {
            const res = await api.post("/auth/login", { email, password });

            if (res.data?.accessToken) {
                localStorage.setItem("accessToken", res.data.accessToken);
                navigate("/");
            } else {
                setError("Invalid response from server.");
            }
        } catch (err) {
            if (err.response) {
                setError(err.response.data?.message || "Login failed. Please check your credentials.");
            } else if (err.request) {
                setError("No response from server. Please try again later.");
            } else {
                setError("An unexpected error occurred.");
            }
        }
    };

    return (
        <div className="auth-container">
            <form onSubmit={handleLogin} className="auth-form">
                <h2 className="auth-title">Login</h2>
                <input
                    className="auth-input"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <input
                    className="auth-input"
                    placeholder="Password"
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <button type="submit" className="auth-button">
                    Login
                </button>

                {error && <p className="auth-error">{error}</p>}
            </form>
        </div>
    );
}
