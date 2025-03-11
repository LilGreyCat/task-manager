import React from "react";
import { Container, Typography } from "@mui/material";
import UserFormView from "./UserFormView";

const AdminView = () => {
  return (
    <Container>
      <Typography variant="h3">Admin Panel</Typography>
      <UserFormView />
    </Container>
  );
};

export default AdminView;
