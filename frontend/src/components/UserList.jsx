import React from "react";
import { List, Typography } from "@mui/material";
import UserItem from "./UserItem";

const UserList = ({ users, onDelete }) => {
  return (
    <>
      <Typography variant="h5" align="center">Users List</Typography>
      <List>
        {users.map(user => (
          <UserItem key={user.id} user={user} onDelete={onDelete} />
        ))}
      </List>
    </>
  );
};

export default UserList;
