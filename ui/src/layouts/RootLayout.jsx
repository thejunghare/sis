import { NavLink, Outlet } from "react-router-dom";

const RootLayout = () => {
  return (
    <div>
      <header className="bg-white shadow">
        <nav className="flex flex-row p-3 items-center justify-between">
          <NavLink to="/">Login</NavLink>
          <NavLink to="/dashboard">Dashboard</NavLink>
        </nav>
      </header>

      <main>
        <Outlet />
      </main>
    </div>
  );
};

export default RootLayout;
