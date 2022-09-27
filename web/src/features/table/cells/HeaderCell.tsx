import React from 'react';

import { LockClosed12Regular as LockIcon } from '@fluentui/react-icons';
import classNames from 'classnames';

import styles from './cell.module.scss';

interface headerCell {
  heading: string;
  className?: string;
  locked?: boolean;
}

function HeaderCell({ heading, className, locked }: headerCell) {
  return (
    <div className={classNames(styles.header, className)}>
      <div className={styles.content}>
        {locked ? (
          <div className="icon small">
            <LockIcon />
          </div>
        ) : (
          ''
        )}
        <div className="text">{heading}</div>
      </div>
    </div>
  );
}
const HeaderCellMemo = React.memo(HeaderCell);
export default HeaderCellMemo;
