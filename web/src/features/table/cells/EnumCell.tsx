import React from 'react';

import { FluentIconsProps } from '@fluentui/react-icons';
import classNames from 'classnames';

import styles from './cell.module.scss';

interface EnumCellProps {
  value: string;
  icon?: React.FC<FluentIconsProps>;
  className?: string;
}

function EnumCell(props: EnumCellProps) {
  return (
    <div className={classNames(styles.cell, styles.alignLeft, styles.EnumCell, props.className)}>
      <div className={styles.current}>
        <>
          {props.icon ? props.icon : ''}
          {props.value}
        </>
      </div>
    </div>
  );
}

const EnumCellMemo = React.memo(EnumCell);
export default EnumCellMemo;
