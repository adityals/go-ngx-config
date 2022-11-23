import type { PropsWithChildren } from 'react';
import NavBar from '../NavBar';

const MainLayout = (props: PropsWithChildren) => {
  return (
    <>
      <NavBar />
      <main>{props.children}</main>
    </>
  );
};

export default MainLayout;
