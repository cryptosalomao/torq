import classNames from 'classnames';
import { Outlet } from 'react-router-dom';

import styles from './default-layout.module.scss';
import LoginLayout from './LoginLayout';

import navStyles from '../features/navigation/nav.module.scss';
import Navigation from '../features/navigation/Navigation';
import { selectHidden } from '../features/navigation/navSlice';
import TopNavigation from '../features/navigation/TopNavigation';
import { useAppSelector } from '../store/hooks';

function DefaultLayout() {
  const xl = <LoginLayout />;
  const hidden = useAppSelector(selectHidden);
  return (
    <div className={classNames(styles.mainContentWrapper, { [navStyles.navCollapsed]: hidden })}>
      <TopNavigation />
      <div className={navStyles.navigationWrapper}>
        <Navigation />
      </div>
      <div className={styles.pageWrapper}>
        {/*<div className="dismiss-navigation-background" onClick={() => dispatch(toggleNav())} />*/}
        <Outlet />
      </div>
    </div>
  );
}

export default DefaultLayout;
