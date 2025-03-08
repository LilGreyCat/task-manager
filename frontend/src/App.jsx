import React, { useEffect, useState } from "react";
import axios from "axios";
import { Container, Typography, List, ListItem, ListItemText } from "@mui/material";

function App() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    axios.get("http://localhost:8080/users")
      .then((response) => {
        setUsers(response.data);
      })
      .catch((error) => {
        console.error("Error fetching users:", error);
      });
  }, []);

  return (
    <Container maxWidth="md">
      <Typography variant="h3" align="center" gutterBottom>
        Task Manager
      </Typography>
      <Typography variant="h5" align="center">
        Users List
      </Typography>
      <List>
        {users.map((user) => (
          <ListItem key={user.id}>
            <ListItemText primary={user.name} secondary={user.email} />
          </ListItem>
        ))}
      </List>
    </Container>
  );
}

export default App;
