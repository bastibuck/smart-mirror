import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";
import { FiMenu } from "react-icons/fi";

const adminPages = [
  { path: "/screens", title: "Screens" },
  { path: "/sections", title: "Sections" },
  { path: "/apps", title: "Apps" },
];

const FloatingMenu = () => {
  const router = useRouter();

  return (
    <div className="absolute right-12 bottom-12">
      <div className="dropdown dropdown-top dropdown-end">
        <label tabIndex={0} className="btn btn-circle m-2 text-2xl">
          <FiMenu />
        </label>

        <ul
          tabIndex={0}
          className="dropdown-content menu shadow bg-base-100  rounded-box w-60 mr-2"
        >
          {adminPages.map((page) => (
            <li key={page.path}>
              <Link href={`/admin${page.path}`}>
                <a
                  className={
                    router.asPath === `/admin${page.path}`
                      ? "active"
                      : undefined
                  }
                >
                  {page.title}
                </a>
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default FloatingMenu;
