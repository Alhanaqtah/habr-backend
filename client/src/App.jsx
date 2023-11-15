import { Route, Routes } from "react-router-dom";

import {
  Articles,
  Followers,
  Following,
  Publications,
  Search,
  User,
} from "./Pages/index.js";

import { Header } from "./Components/index.js";

import userIcon from "./assets/user.png";
import searchIcon from "./assets/search.png";
import avatarPlaceholder from "./assets/avatarPlaceholder.jpg";

export default function App() {
  const headerItems = [
    { id: 1, linkTo: "articles", name: "Articles", active: true },
    { id: 2, linkTo: "feed/articles", name: "Personal articles", active: true },
  ];

  return (
    <>
      <Header
        headerItems={headerItems}
        userIcon={userIcon}
        searchIcon={searchIcon}
      />

      <Routes>
        <Route path="articles" element={<Articles />} />
        <Route path="feed/articles" element={<Articles />} />
        <Route path="search" element={<Search />} />
        <Route path="users/user/publications" element={<Publications />} />
        <Route path="users/user/followers" element={<Followers />} />
        <Route path="users/user/following" element={<Following />} />
        <Route
          path="users/user"
          element={<User userIcon={avatarPlaceholder} />}
        />
      </Routes>
    </>
  );
}
