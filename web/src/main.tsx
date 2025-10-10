import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router";
import {
    LoginPage,
    CreateAccountPage,
    ForgotPasswordPage,
    ResetPasswordPage,
    IndexPage,
    ErrorPage,
    NotFoundPage,
    LogoutPage
} from "./pages";
import "./index.css";

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        <BrowserRouter>
            <Routes>
                <Route index path="/" element={<IndexPage />} />
                <Route path="/login" element={<LoginPage />} />
                <Route path="/create-account" element={<CreateAccountPage />} />
                <Route
                    path="/forgot-password"
                    element={<ForgotPasswordPage />}
                />
                <Route path="/reset-password" element={<ResetPasswordPage />} />
                <Route path="/error" element={<ErrorPage />} />
                <Route path="/logout" element={<LogoutPage />} />
                <Route path="*" element={<NotFoundPage />} />
            </Routes>
        </BrowserRouter>
    </StrictMode>
);
