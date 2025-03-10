import React from "react";
import { TextField, Button, Container, Typography } from "@mui/material";
import { useUserForm } from "../hooks/useUserForm";

const UserFormView = ({ onUserAdded }) => {
  const { formData, handleChange, handleSubmit } = useUserForm(onUserAdded);

  return (
    <Container>
      <Typography variant="h5">Create User</Typography>
      <form onSubmit={handleSubmit}>
        <TextField fullWidth name="name" label="Name" value={formData.name} onChange={handleChange} margin="normal" />
        <TextField fullWidth name="email" label="Email" value={formData.email} onChange={handleChange} margin="normal" />
        <TextField fullWidth type="password" name="password" label="Password" value={formData.password} onChange={handleChange} margin="normal" />
        <Button type="submit" variant="contained" color="primary">Create User</Button>
      </form>
    </Container>
  );
};

export default UserFormView;
