import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api/axios";

export default function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");   // error message state
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();
        setError(""); // reset error on new submit
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
                // API responded with a status code outside 2xx
                setError(err.response.data?.message || "Login failed. Please check your credentials.");
            } else if (err.request) {
                // Request was made but no response
                setError("No response from server. Please try again later.");
            } else {
                // Other errors (setup, parsing, etc.)
                setError("An unexpected error occurred.");
            }
        }
    };

    return (
        <form onSubmit={handleLogin} className="form">
            <h2>Login</h2>
            <input
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
            />
            <input
                placeholder="Password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />
            <button type="submit">Login</button>

            {error && <p style={{ color: "red" }}>{error}</p>}
        </form>
    );
}
