import React from "react";
import { Container, Typography, Button } from "@mui/material";
import { Link } from "react-router-dom";

const HomeView = () => {
  return (
    <Container>
      <Typography variant="h3">Welcome to Task Manager</Typography>
      <Button variant="contained" component={Link} to="/admin" color="primary">
        Go to Admin Panel
      </Button>
    </Container>
  );
};

export default HomeView;
