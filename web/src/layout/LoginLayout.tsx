import classNames from 'classnames';
import { Outlet } from 'react-router-dom';

import './login_layout.scss';

function LoginLayout() {
  return (
    <div className={classNames('login-layout')}>
      <Outlet />
    </div>
  );
}

export default LoginLayout;
