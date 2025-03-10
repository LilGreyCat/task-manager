import React from "react";
import { ListItem, ListItemText, Button } from "@mui/material";

const UserItem = ({ user, onDelete }) => {
  return (
    <ListItem sx={{ display: "flex", justifyContent: "space-between" }}>
      <ListItemText primary={user.name} secondary={user.email} />
      <Button variant="contained" color="error" onClick={() => onDelete(user.id)}>Delete</Button>
    </ListItem>
  );
};

export default UserItem;
