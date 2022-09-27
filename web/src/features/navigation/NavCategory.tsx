import React from 'react';

import { LineHorizontal120Regular as CollapseIcon, ChevronDown20Regular as ExpandIcon } from '@fluentui/react-icons';
import classNames from 'classnames';

import styles from './nav.module.scss';

function NavCategory(props: { text: string; collapsed?: boolean; children: React.ReactNode }) {
  const icon = props.collapsed ? <ExpandIcon /> : <CollapseIcon />;
  return (
    <div className={classNames(styles.navCategory)}>
      <div
        className={classNames(
          styles.NavCategoryTitleContainer,
          styles.navCollapsedCategoryTitle,
          styles.navCategoryTitle
        )}
      >
        <CollapseIcon />
      </div>
      <div className={classNames(styles.NavCategoryTitleContainer)}>
        <div className={classNames(styles.navCategoryTitle)}>{props.text}</div>
        <div className={classNames(styles.icon)}>{icon}</div>
      </div>
      {!props.collapsed && <div>{props.children}</div>}
    </div>
  );
}

export default React.memo(NavCategory);
