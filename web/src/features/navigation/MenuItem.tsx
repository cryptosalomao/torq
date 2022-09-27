import React from 'react';

import classNames from 'classnames';
import { NavLink } from 'react-router-dom';

import styles from './nav.module.scss';

function MenuItem(props: { text: string; icon?: any; routeTo: string }) {
  const linkClasses = (p: { isActive: boolean }): string => {
    return classNames(styles.title, { [styles.selected]: p.isActive });
  };

  return (
    <div className={classNames(styles.item)}>
      <div className={classNames(styles.contentWrapper)}>
        <NavLink to={props.routeTo} className={linkClasses}>
          <div className={styles.icon}>{props.icon}</div>
          <div className={classNames(styles.text)}>{props.text}</div>
        </NavLink>
      </div>
    </div>
  );
}

export default React.memo(MenuItem);
