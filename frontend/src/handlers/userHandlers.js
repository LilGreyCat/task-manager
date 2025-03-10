import axios from "axios";

const API_URL = "http://localhost:8080/users";

// Fetch all users
export const fetchUsers = async (setUsers) => {
  try {
    const response = await axios.get(API_URL);
    setUsers(response.data);
  } catch (error) {
    console.error("Error fetching users:", error);
  }
};

// Delete a user
export const deleteUser = async (id, setUsers) => {
  try {
    console.log(`Deleting user with ID: ${id}`);  // Debugging log
    await axios.delete(`${API_URL}/${id}`);
    setUsers(prevUsers => prevUsers.filter(user => user.id !== id));
  } catch (error) {
    console.error("Error deleting user:", error);
  }
};

// Create a new user
export const createUser = async (userData, onUserAdded) => {
  try {
    await axios.post(API_URL, userData);
    onUserAdded(); // Refresh user list
  } catch (error) {
    console.error("Error creating user:", error);
  }
};
