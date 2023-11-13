import { Route, Routes } from "react-router-dom";

import {
  Articles,
  Followers,
  Following,
  Publications,
  Search,
  User,
} from "./Pages/index.js";

import { Header } from "./Components";

import userIcon from "./assets/user.png";

export default function App() {
  return (
    <>
      <Header userIcon={userIcon} />

      <Routes>
        <Route path="articles" element={<Articles />} />
        <Route path="feed/articles" element={<Articles />} />
        <Route path="search" element={<Search />} />
        <Route path="users/publications" element={<Publications />} />
        <Route path="users/followers" element={<Followers />} />
        <Route path="users/following" element={<Following />} />
        <Route path="users/user" element={<User />} />
      </Routes>
    </>
  );
}
