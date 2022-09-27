import { VirtualNetwork20Regular as ChannelsIcon, Iot20Regular as PeersIcon } from '@fluentui/react-icons';
import classNames from 'classnames';

import styles from './group-section.module.scss';

type GroupPopoverProps = {
  groupBy: string;
  groupByHandler: (groupBy: string) => void;
};

const GroupBySection = (props: GroupPopoverProps) => {
  return (
    <div className={styles.groupRowWrapper}>
      <button
        className={classNames(styles.groupRow, {
          [styles.groupRowSelected]: props.groupBy == 'channels',
        })}
        onClick={() => {
          props.groupByHandler('channels');
        }}
      >
        <div className="icon">
          <ChannelsIcon />
        </div>
        <div>Channels</div>
      </button>
      <button
        className={classNames(styles.groupRow, {
          [styles.groupRowSelected]: props.groupBy == 'peers',
        })}
        onClick={() => {
          props.groupByHandler('peers');
        }}
      >
        <div className="icon">
          <PeersIcon />
        </div>
        <div>Peers</div>
      </button>
      {/*<button*/}
      {/*  className={classNames(styles.groupRow, {[styles.groupRowSelected]: (props.groupBy == 'tags')})}*/}
      {/*  onClick={() =>{*/}
      {/*    dispatch(updateGroupBy({groupBy: 'tags'}))*/}
      {/*  }}*/}
      {/*>*/}
      {/*  <div className="icon">*/}
      {/*    <TagsIcon/>*/}
      {/*  </div>*/}
      {/*  <div>Tags</div>*/}
      {/*</button>*/}
    </div>
  );
};

export default GroupBySection;
