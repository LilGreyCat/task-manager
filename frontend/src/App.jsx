import React, { useEffect, useState } from "react";
import { Container, Typography } from "@mui/material";
import UserFormView from "./views/UserFormView";
import UserList from "./components/UserList";
import { fetchUsers, deleteUser } from "./handlers/userHandlers";

function App() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    fetchUsers(setUsers);
  }, []);

  return (
    <Container maxWidth="md">
      <Typography variant="h3" align="center" gutterBottom>
        Task Manager - Users
      </Typography>

      <UserFormView onUserAdded={() => fetchUsers(setUsers)} />
      <UserList users={users} onDelete={(id) => deleteUser(id, setUsers)} />
    </Container>
  );
}

export default App;
