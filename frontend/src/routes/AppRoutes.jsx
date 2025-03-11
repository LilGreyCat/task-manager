import React from "react";
import { Routes, Route } from "react-router-dom";
import HomeView from "../views/HomeView";
import LoginView from "../views/LoginView";
import AdminView from "../views/AdminView";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<HomeView />} />
      <Route path="/login" element={<LoginView />} />
      <Route path="/admin" element={<AdminView />} />
    </Routes>
  );
};

export default AppRoutes;
