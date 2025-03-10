import { useState } from "react";
import { createUser } from "../handlers/userHandlers";

export const useUserForm = (onUserAdded) => {
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    password: "",
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    await createUser(formData, onUserAdded);
    setFormData({ name: "", email: "", password: "" });
  };

  return { formData, handleChange, handleSubmit };
};
