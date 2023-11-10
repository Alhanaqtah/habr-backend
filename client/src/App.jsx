import { Route, Routes } from "react-router-dom";

import {
  Articles,
  Followers,
  Following,
  Publications,
  Search,
  User,
  Users,
} from "./Pages/index.js";

import { Header } from "./Components";

export default function App() {
  return (
    <>
      <Header />

      <Routes>
        <Route path="articles" element={<Articles />} />
        <Route path="feed/articles" element={<Articles />} />
        <Route path="users/followers" element={<Followers />} />
        <Route path="users/following" element={<Following />} />
        <Route path="publications" element={<Publications />} />
        <Route path="search" element={<Search />} />
        <Route path="users/user" element={<User />} />
        <Route path="users" element={<Users />} />
      </Routes>
    </>
  );
}
