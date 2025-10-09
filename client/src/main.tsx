import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router";
import {
    LoginPage,
    CreateAccountPage,
    ForgotPasswordPage,
    ResetPasswordPage,
    IndexPage
} from "./pages";
import "./index.css";

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        <BrowserRouter>
            <Routes>
                <Route index path="/" element={<IndexPage />} />
                <Route index path="/login" element={<LoginPage />} />
                <Route path="/create-account" element={<CreateAccountPage />} />
                <Route
                    path="/forgot-password"
                    element={<ForgotPasswordPage />}
                />
                <Route path="/reset-password" element={<ResetPasswordPage />} />
            </Routes>
        </BrowserRouter>
    </StrictMode>
);
