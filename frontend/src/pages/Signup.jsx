import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api/axios";

export default function Signup() {
    const navigate = useNavigate();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [error, setError] = useState("");
    const [success, setSuccess] = useState("");

    const validateEmail = (email) => {
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
    };

    const handleSignup = async (e) => {
        e.preventDefault();
        setError("");
        setSuccess("");

        if (!validateEmail(email)) {
            setError("Please enter a valid email address.");
            return;
        }

        if (password !== confirmPassword) {
            setError("Passwords do not match.");
            return;
        }

        try {
            await api.post("/auth/signup", { email, password });
            setSuccess("Signup successful. Please login.");
            navigate("/login");
        } catch (err) {
            if (err.response) {
                setError(err.response.data?.message || "Signup failed. Please try again.");
            } else if (err.request) {
                setError("No response from server. Please try again later.");
            } else {
                setError("An unexpected error occurred.");
            }
        }
    };

    return (
        <div className="auth-container">
            <form onSubmit={handleSignup} className="auth-form">
                <h2 className="auth-title">Signup</h2>
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
                <input
                    className="auth-input"
                    placeholder="Confirm Password"
                    type="password"
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                />
                <button type="submit" className="auth-button">
                    Signup
                </button>

                {error && <p className="auth-error">{error}</p>}
                {success && <p className="auth-success">{success}</p>}
            </form>
        </div>
    );
}
