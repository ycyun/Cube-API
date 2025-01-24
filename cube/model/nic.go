package model

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/ycyun/Cube-API/utils"
	"net/http"
	"os/exec"
	"reflect"
	"sync"
	"time"
)

// ip -d -j a show
type NICStatus struct {
	Ifindex        int      `json:"ifindex"`
	Ifname         string   `json:"ifname"`
	Flags          []string `json:"flags"`
	Mtu            int      `json:"mtu"`
	Qdisc          string   `json:"qdisc"`
	Operstate      string   `json:"operstate"`
	Group          string   `json:"group"`
	Txqlen         int      `json:"txqlen"`
	LinkType       string   `json:"link_type"`
	Address        string   `json:"address"`
	Broadcast      string   `json:"broadcast"`
	Promiscuity    int      `json:"promiscuity"`
	Allmulti       int      `json:"allmulti"`
	MinMtu         int      `json:"min_mtu"`
	MaxMtu         int      `json:"max_mtu"`
	NumTxQueues    int      `json:"num_tx_queues"`
	NumRxQueues    int      `json:"num_rx_queues"`
	GsoMaxSize     int      `json:"gso_max_size"`
	GsoMaxSegs     int      `json:"gso_max_segs"`
	TsoMaxSize     int      `json:"tso_max_size"`
	TsoMaxSegs     int      `json:"tso_max_segs"`
	GroMaxSize     int      `json:"gro_max_size"`
	GsoIpv4MaxSize int      `json:"gso_ipv4_max_size"`
	GroIpv4MaxSize int      `json:"gro_ipv4_max_size"`
	AddrInfo       []struct {
		Family            string `json:"family"`
		Local             string `json:"local"`
		Prefixlen         int    `json:"prefixlen"`
		Scope             string `json:"scope"`
		Label             string `json:"label,omitempty"`
		ValidLifeTime     int64  `json:"valid_life_time"`
		PreferredLifeTime int64  `json:"preferred_life_time"`
		Broadcast         string `json:"broadcast,omitempty"`
		Noprefixroute     bool   `json:"noprefixroute,omitempty"`
	} `json:"addr_info"`
	Parentbus    string        `json:"parentbus,omitempty"`
	Parentdev    string        `json:"parentdev,omitempty"`
	Altnames     []string      `json:"altnames,omitempty"`
	PhysPortName string        `json:"phys_port_name,omitempty"`
	PhysSwitchId string        `json:"phys_switch_id,omitempty"`
	VfinfoList   []interface{} `json:"vfinfo_list,omitempty"`
	Master       string        `json:"master,omitempty"`
	Linkinfo     struct {
		InfoSlaveKind string `json:"info_slave_kind,omitempty"`
		InfoSlaveData struct {
			State             string  `json:"state"`
			MiiStatus         string  `json:"mii_status,omitempty"`
			LinkFailureCount  int     `json:"link_failure_count,omitempty"`
			PermHwaddr        string  `json:"perm_hwaddr,omitempty"`
			QueueId           int     `json:"queue_id,omitempty"`
			Prio              int     `json:"prio,omitempty"`
			Priority          int     `json:"priority,omitempty"`
			Cost              int     `json:"cost,omitempty"`
			Hairpin           bool    `json:"hairpin,omitempty"`
			Guard             bool    `json:"guard,omitempty"`
			RootBlock         bool    `json:"root_block,omitempty"`
			Fastleave         bool    `json:"fastleave,omitempty"`
			Learning          bool    `json:"learning,omitempty"`
			Flood             bool    `json:"flood,omitempty"`
			Id                string  `json:"id,omitempty"`
			No                string  `json:"no,omitempty"`
			DesignatedPort    int     `json:"designated_port,omitempty"`
			DesignatedCost    int     `json:"designated_cost,omitempty"`
			BridgeId          string  `json:"bridge_id,omitempty"`
			RootId            string  `json:"root_id,omitempty"`
			HoldTimer         float64 `json:"hold_timer,omitempty"`
			MessageAgeTimer   float64 `json:"message_age_timer,omitempty"`
			ForwardDelayTimer float64 `json:"forward_delay_timer,omitempty"`
			TopologyChangeAck int     `json:"topology_change_ack,omitempty"`
			ConfigPending     int     `json:"config_pending,omitempty"`
			ProxyArp          bool    `json:"proxy_arp,omitempty"`
			ProxyArpWifi      bool    `json:"proxy_arp_wifi,omitempty"`
			MulticastRouter   int     `json:"multicast_router,omitempty"`
			McastFlood        bool    `json:"mcast_flood,omitempty"`
			BcastFlood        bool    `json:"bcast_flood,omitempty"`
			McastToUnicast    bool    `json:"mcast_to_unicast,omitempty"`
			NeighSuppress     bool    `json:"neigh_suppress,omitempty"`
			GroupFwdMask      string  `json:"group_fwd_mask,omitempty"`
			GroupFwdMaskStr   string  `json:"group_fwd_mask_str,omitempty"`
			VlanTunnel        bool    `json:"vlan_tunnel,omitempty"`
			Isolated          bool    `json:"isolated,omitempty"`
			Locked            bool    `json:"locked,omitempty"`
			Mab               bool    `json:"mab,omitempty"`
		} `json:"info_slave_data,omitempty"`
		InfoKind string `json:"info_kind,omitempty"`
		InfoData struct {
			Mode                    string      `json:"mode,omitempty"`
			ActiveSlave             string      `json:"active_slave,omitempty"`
			Miimon                  int         `json:"miimon,omitempty"`
			Updelay                 int         `json:"updelay,omitempty"`
			Downdelay               int         `json:"downdelay,omitempty"`
			PeerNotifyDelay         int         `json:"peer_notify_delay,omitempty"`
			UseCarrier              int         `json:"use_carrier,omitempty"`
			ArpInterval             int         `json:"arp_interval,omitempty"`
			ArpMissedMax            int         `json:"arp_missed_max,omitempty"`
			ArpValidate             interface{} `json:"arp_validate"`
			ArpAllTargets           string      `json:"arp_all_targets,omitempty"`
			PrimaryReselect         string      `json:"primary_reselect,omitempty"`
			FailOverMac             string      `json:"fail_over_mac,omitempty"`
			XmitHashPolicy          string      `json:"xmit_hash_policy,omitempty"`
			ResendIgmp              int         `json:"resend_igmp,omitempty"`
			NumPeerNotif            int         `json:"num_peer_notif,omitempty"`
			AllSlavesActive         int         `json:"all_slaves_active,omitempty"`
			MinLinks                int         `json:"min_links,omitempty"`
			LpInterval              int         `json:"lp_interval,omitempty"`
			PacketsPerSlave         int         `json:"packets_per_slave,omitempty"`
			AdLacpActive            string      `json:"ad_lacp_active,omitempty"`
			AdLacpRate              string      `json:"ad_lacp_rate,omitempty"`
			AdSelect                string      `json:"ad_select,omitempty"`
			TlbDynamicLb            int         `json:"tlb_dynamic_lb,omitempty"`
			ForwardDelay            int         `json:"forward_delay,omitempty"`
			HelloTime               int         `json:"hello_time,omitempty"`
			MaxAge                  int         `json:"max_age,omitempty"`
			AgeingTime              int         `json:"ageing_time,omitempty"`
			StpState                int         `json:"stp_state,omitempty"`
			Priority                int         `json:"priority,omitempty"`
			VlanFiltering           int         `json:"vlan_filtering,omitempty"`
			VlanProtocol            string      `json:"vlan_protocol,omitempty"`
			BridgeId                string      `json:"bridge_id,omitempty"`
			RootId                  string      `json:"root_id,omitempty"`
			RootPort                int         `json:"root_port,omitempty"`
			RootPathCost            int         `json:"root_path_cost,omitempty"`
			TopologyChange          int         `json:"topology_change,omitempty"`
			TopologyChangeDetected  int         `json:"topology_change_detected,omitempty"`
			HelloTimer              float64     `json:"hello_timer,omitempty"`
			TcnTimer                float64     `json:"tcn_timer,omitempty"`
			TopologyChangeTimer     float64     `json:"topology_change_timer,omitempty"`
			GcTimer                 float64     `json:"gc_timer,omitempty"`
			VlanDefaultPvid         int         `json:"vlan_default_pvid,omitempty"`
			VlanStatsEnabled        int         `json:"vlan_stats_enabled,omitempty"`
			VlanStatsPerPort        int         `json:"vlan_stats_per_port,omitempty"`
			GroupFwdMask            string      `json:"group_fwd_mask,omitempty"`
			GroupAddr               string      `json:"group_addr,omitempty"`
			McastSnooping           int         `json:"mcast_snooping,omitempty"`
			NoLinklocalLearn        int         `json:"no_linklocal_learn,omitempty"`
			McastVlanSnooping       int         `json:"mcast_vlan_snooping,omitempty"`
			McastRouter             int         `json:"mcast_router,omitempty"`
			McastQueryUseIfaddr     int         `json:"mcast_query_use_ifaddr,omitempty"`
			McastQuerier            int         `json:"mcast_querier,omitempty"`
			McastHashElasticity     int         `json:"mcast_hash_elasticity,omitempty"`
			McastHashMax            int         `json:"mcast_hash_max,omitempty"`
			McastLastMemberCnt      int         `json:"mcast_last_member_cnt,omitempty"`
			McastStartupQueryCnt    int         `json:"mcast_startup_query_cnt,omitempty"`
			McastLastMemberIntvl    int         `json:"mcast_last_member_intvl,omitempty"`
			McastMembershipIntvl    int         `json:"mcast_membership_intvl,omitempty"`
			McastQuerierIntvl       int         `json:"mcast_querier_intvl,omitempty"`
			McastQueryIntvl         int         `json:"mcast_query_intvl,omitempty"`
			McastQueryResponseIntvl int         `json:"mcast_query_response_intvl,omitempty"`
			McastStartupQueryIntvl  int         `json:"mcast_startup_query_intvl,omitempty"`
			McastStatsEnabled       int         `json:"mcast_stats_enabled,omitempty"`
			McastIgmpVersion        int         `json:"mcast_igmp_version,omitempty"`
			McastMldVersion         int         `json:"mcast_mld_version,omitempty"`
			NfCallIptables          int         `json:"nf_call_iptables,omitempty"`
			NfCallIp6Tables         int         `json:"nf_call_ip6tables,omitempty"`
			NfCallArptables         int         `json:"nf_call_arptables,omitempty"`
			Type                    string      `json:"type,omitempty"`
			Pi                      bool        `json:"pi,omitempty"`
			VnetHdr                 bool        `json:"vnet_hdr,omitempty"`
			MultiQueue              bool        `json:"multi_queue,omitempty"`
			Persist                 bool        `json:"persist,omitempty"`
		} `json:"info_data,omitempty"`
	} `json:"linkinfo,omitempty"`
} // @name NICStatus

type TypeNICStatus struct {
	NICs        []*NICStatus `json:"nics"`
	RefreshTime time.Time
}
type TypeNicStatus2 struct {
	Ifindex     int      `json:"ifindex"`
	Ifname      string   `json:"ifname"`
	Flags       []string `json:"flags"`
	Mtu         int      `json:"mtu"`
	Qdisc       string   `json:"qdisc"`
	Operstate   string   `json:"operstate"`
	Group       string   `json:"group"`
	Txqlen      int      `json:"txqlen"`
	LinkType    string   `json:"link_type"`
	Address     string   `json:"address"`
	Broadcast   string   `json:"broadcast"`
	Promiscuity int      `json:"promiscuity"`
	MinMtu      int      `json:"min_mtu"`
	MaxMtu      int      `json:"max_mtu"`
	NumTxQueues int      `json:"num_tx_queues"`
	NumRxQueues int      `json:"num_rx_queues"`
	GsoMaxSize  int      `json:"gso_max_size"`
	GsoMaxSegs  int      `json:"gso_max_segs"`
	TsoMaxSize  int      `json:"tso_max_size"`
	TsoMaxSegs  int      `json:"tso_max_segs"`
	GroMaxSize  int      `json:"gro_max_size"`
	AddrInfo    []struct {
		Family            string `json:"family"`
		Local             string `json:"local"`
		Prefixlen         int    `json:"prefixlen"`
		Scope             string `json:"scope"`
		Label             string `json:"label,omitempty"`
		ValidLifeTime     int64  `json:"valid_life_time"`
		PreferredLifeTime int64  `json:"preferred_life_time"`
		Broadcast         string `json:"broadcast,omitempty"`
		Noprefixroute     bool   `json:"noprefixroute,omitempty"`
	} `json:"addr_info"`
	Master   string `json:"master,omitempty"`
	Linkinfo struct {
		InfoSlaveKind string `json:"info_slave_kind,omitempty"`
		InfoSlaveData struct {
			State             string  `json:"state"`
			Priority          int     `json:"priority"`
			Cost              int     `json:"cost"`
			Hairpin           bool    `json:"hairpin"`
			Guard             bool    `json:"guard"`
			RootBlock         bool    `json:"root_block"`
			Fastleave         bool    `json:"fastleave"`
			Learning          bool    `json:"learning"`
			Flood             bool    `json:"flood"`
			Id                string  `json:"id"`
			No                string  `json:"no"`
			DesignatedPort    int     `json:"designated_port"`
			DesignatedCost    int     `json:"designated_cost"`
			BridgeId          string  `json:"bridge_id"`
			RootId            string  `json:"root_id"`
			HoldTimer         float64 `json:"hold_timer"`
			MessageAgeTimer   float64 `json:"message_age_timer"`
			ForwardDelayTimer float64 `json:"forward_delay_timer"`
			TopologyChangeAck int     `json:"topology_change_ack"`
			ConfigPending     int     `json:"config_pending"`
			ProxyArp          bool    `json:"proxy_arp"`
			ProxyArpWifi      bool    `json:"proxy_arp_wifi"`
			MulticastRouter   int     `json:"multicast_router"`
			McastFlood        bool    `json:"mcast_flood"`
			BcastFlood        bool    `json:"bcast_flood"`
			McastToUnicast    bool    `json:"mcast_to_unicast"`
			NeighSuppress     bool    `json:"neigh_suppress"`
			GroupFwdMask      string  `json:"group_fwd_mask"`
			GroupFwdMaskStr   string  `json:"group_fwd_mask_str"`
			VlanTunnel        bool    `json:"vlan_tunnel"`
			Isolated          bool    `json:"isolated"`
			Locked            bool    `json:"locked"`
		} `json:"info_slave_data,omitempty"`
		InfoKind string `json:"info_kind,omitempty"`
		InfoData struct {
			ForwardDelay            int      `json:"forward_delay,omitempty"`
			HelloTime               int      `json:"hello_time,omitempty"`
			MaxAge                  int      `json:"max_age,omitempty"`
			AgeingTime              int      `json:"ageing_time,omitempty"`
			StpState                int      `json:"stp_state,omitempty"`
			Priority                int      `json:"priority,omitempty"`
			VlanFiltering           int      `json:"vlan_filtering,omitempty"`
			VlanProtocol            string   `json:"vlan_protocol,omitempty"`
			BridgeId                string   `json:"bridge_id,omitempty"`
			RootId                  string   `json:"root_id,omitempty"`
			RootPort                int      `json:"root_port,omitempty"`
			RootPathCost            int      `json:"root_path_cost,omitempty"`
			TopologyChange          int      `json:"topology_change,omitempty"`
			TopologyChangeDetected  int      `json:"topology_change_detected,omitempty"`
			HelloTimer              float64  `json:"hello_timer,omitempty"`
			TcnTimer                float64  `json:"tcn_timer,omitempty"`
			TopologyChangeTimer     float64  `json:"topology_change_timer,omitempty"`
			GcTimer                 float64  `json:"gc_timer,omitempty"`
			VlanDefaultPvid         int      `json:"vlan_default_pvid,omitempty"`
			VlanStatsEnabled        int      `json:"vlan_stats_enabled,omitempty"`
			VlanStatsPerPort        int      `json:"vlan_stats_per_port,omitempty"`
			GroupFwdMask            string   `json:"group_fwd_mask,omitempty"`
			GroupAddr               string   `json:"group_addr,omitempty"`
			McastSnooping           int      `json:"mcast_snooping,omitempty"`
			NoLinklocalLearn        int      `json:"no_linklocal_learn,omitempty"`
			McastVlanSnooping       int      `json:"mcast_vlan_snooping,omitempty"`
			McastRouter             int      `json:"mcast_router,omitempty"`
			McastQueryUseIfaddr     int      `json:"mcast_query_use_ifaddr,omitempty"`
			McastQuerier            int      `json:"mcast_querier,omitempty"`
			McastHashElasticity     int      `json:"mcast_hash_elasticity,omitempty"`
			McastHashMax            int      `json:"mcast_hash_max,omitempty"`
			McastLastMemberCnt      int      `json:"mcast_last_member_cnt,omitempty"`
			McastStartupQueryCnt    int      `json:"mcast_startup_query_cnt,omitempty"`
			McastLastMemberIntvl    int      `json:"mcast_last_member_intvl,omitempty"`
			McastMembershipIntvl    int      `json:"mcast_membership_intvl,omitempty"`
			McastQuerierIntvl       int      `json:"mcast_querier_intvl,omitempty"`
			McastQueryIntvl         int      `json:"mcast_query_intvl,omitempty"`
			McastQueryResponseIntvl int      `json:"mcast_query_response_intvl,omitempty"`
			McastStartupQueryIntvl  int      `json:"mcast_startup_query_intvl,omitempty"`
			McastStatsEnabled       int      `json:"mcast_stats_enabled,omitempty"`
			McastIgmpVersion        int      `json:"mcast_igmp_version,omitempty"`
			McastMldVersion         int      `json:"mcast_mld_version,omitempty"`
			NfCallIptables          int      `json:"nf_call_iptables,omitempty"`
			NfCallIp6Tables         int      `json:"nf_call_ip6tables,omitempty"`
			NfCallArptables         int      `json:"nf_call_arptables,omitempty"`
			Type                    string   `json:"type,omitempty"`
			Pi                      bool     `json:"pi,omitempty"`
			VnetHdr                 bool     `json:"vnet_hdr,omitempty"`
			MultiQueue              bool     `json:"multi_queue,omitempty"`
			Persist                 bool     `json:"persist,omitempty"`
			Protocol                string   `json:"protocol,omitempty"`
			Id                      int      `json:"id,omitempty"`
			Flags                   []string `json:"flags,omitempty"`
		} `json:"info_data,omitempty"`
	} `json:"linkinfo,omitempty"`
	PhysPortId string        `json:"phys_port_id,omitempty"`
	Parentbus  string        `json:"parentbus,omitempty"`
	Parentdev  string        `json:"parentdev,omitempty"`
	VfinfoList []interface{} `json:"vfinfo_list,omitempty"`
	Altnames   []string      `json:"altnames,omitempty"`
	Link       string        `json:"link,omitempty"`
}

var lockNICStatus sync.Once

var _NICStatus *TypeNICStatus

func NIC() *TypeNICStatus {
	if _NICStatus == nil {
		lockNICStatus.Do(
			func() {
				if gin.IsDebugging() {
					fmt.Println("Creating ", reflect.TypeOf(_NICStatus), " now.")
				}
				_NICStatus = &TypeNICStatus{}
			})
	} else {
		if gin.IsDebugging() {
			fmt.Println("get old ", reflect.TypeOf(_NICStatus), " instance.")
		}
	}
	return _NICStatus
}

func (nic *TypeNICStatus) Update() {
	var NICs []*NICStatus
	if gin.Mode() == gin.ReleaseMode {
		var stdout []byte

		cmd := exec.Command("ip", "-d", "-j", "address", "show")
		stdout, _ = cmd.CombinedOutput()

		if err := json.Unmarshal(stdout, &NICs); err != nil && gin.IsDebugging() {
			utils.FancyHandleError(err)

		}
	} else {

		stdout := []byte("[{\"ifindex\":1,\"ifname\":\"lo\",\"flags\":[\"LOOPBACK\",\"UP\",\"LOWER_UP\"],\"mtu\":65536,\"qdisc\":\"noqueue\",\"operstate\":\"UNKNOWN\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"loopback\",\"address\":\"00:00:00:00:00:00\",\"broadcast\":\"00:00:00:00:00:00\",\"promiscuity\":0,\"allmulti\":0,\"min_mtu\":0,\"max_mtu\":0,\"num_tx_queues\":1,\"num_rx_queues\":1,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":524280,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"addr_info\":[{\"family\":\"inet\",\"local\":\"127.0.0.1\",\"prefixlen\":8,\"scope\":\"host\",\"label\":\"lo\",\"valid_life_time\":4294967295,\"preferred_life_time\":4294967295},{\"family\":\"inet6\",\"local\":\"::1\",\"prefixlen\":128,\"scope\":\"host\",\"valid_life_time\":4294967295,\"preferred_life_time\":4294967295}]},{\"ifindex\":2,\"ifname\":\"eno8303\",\"flags\":[\"NO-CARRIER\",\"BROADCAST\",\"MULTICAST\",\"UP\"],\"mtu\":1500,\"qdisc\":\"mq\",\"operstate\":\"DOWN\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"ether\",\"address\":\"c4:cb:e1:bf:c2:98\",\"broadcast\":\"ff:ff:ff:ff:ff:ff\",\"promiscuity\":0,\"allmulti\":0,\"min_mtu\":60,\"max_mtu\":9000,\"num_tx_queues\":5,\"num_rx_queues\":5,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":65536,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"parentbus\":\"pci\",\"parentdev\":\"0000:04:00.0\",\"altnames\":[\"enp4s0f0\"],\"addr_info\":[]},{\"ifindex\":3,\"ifname\":\"eno8403\",\"flags\":[\"NO-CARRIER\",\"BROADCAST\",\"MULTICAST\",\"UP\"],\"mtu\":1500,\"qdisc\":\"mq\",\"operstate\":\"DOWN\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"ether\",\"address\":\"c4:cb:e1:bf:c2:99\",\"broadcast\":\"ff:ff:ff:ff:ff:ff\",\"promiscuity\":0,\"allmulti\":0,\"min_mtu\":60,\"max_mtu\":9000,\"num_tx_queues\":5,\"num_rx_queues\":5,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":65536,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"parentbus\":\"pci\",\"parentdev\":\"0000:04:00.1\",\"altnames\":[\"enp4s0f1\"],\"addr_info\":[]},{\"ifindex\":4,\"ifname\":\"eno12399np0\",\"flags\":[\"NO-CARRIER\",\"BROADCAST\",\"MULTICAST\",\"UP\"],\"mtu\":1500,\"qdisc\":\"mq\",\"operstate\":\"DOWN\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"ether\",\"address\":\"d4:04:e6:6f:89:f0\",\"broadcast\":\"ff:ff:ff:ff:ff:ff\",\"promiscuity\":0,\"allmulti\":0,\"min_mtu\":60,\"max_mtu\":9600,\"num_tx_queues\":74,\"num_rx_queues\":74,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":524280,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"phys_port_name\":\"p0\",\"phys_switch_id\":\"f0896ffeffe604d4\",\"parentbus\":\"pci\",\"parentdev\":\"0000:31:00.0\",\"vfinfo_list\":[],\"altnames\":[\"enp49s0f0np0\"],\"addr_info\":[]},{\"ifindex\":5,\"ifname\":\"eno12409np1\",\"flags\":[\"NO-CARRIER\",\"BROADCAST\",\"MULTICAST\",\"UP\"],\"mtu\":1500,\"qdisc\":\"mq\",\"operstate\":\"DOWN\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"ether\",\"address\":\"d4:04:e6:6f:89:f1\",\"broadcast\":\"ff:ff:ff:ff:ff:ff\",\"promiscuity\":0,\"allmulti\":0,\"min_mtu\":60,\"max_mtu\":9600,\"num_tx_queues\":74,\"num_rx_queues\":74,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":524280,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"phys_port_name\":\"p1\",\"phys_switch_id\":\"f0896ffeffe604d4\",\"parentbus\":\"pci\",\"parentdev\":\"0000:31:00.1\",\"vfinfo_list\":[],\"altnames\":[\"enp49s0f1np1\"],\"addr_info\":[]},{\"ifindex\":6,\"ifname\":\"ens1f0np0\",\"flags\":[\"BROADCAST\",\"MULTICAST\",\"SLAVE\",\"UP\",\"LOWER_UP\"],\"mtu\":1500,\"qdisc\":\"mq\",\"master\":\"bond0\",\"operstate\":\"UP\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"ether\",\"address\":\"00:62:0b:e7:99:50\",\"broadcast\":\"ff:ff:ff:ff:ff:ff\",\"promiscuity\":1,\"allmulti\":1,\"min_mtu\":60,\"max_mtu\":9600,\"linkinfo\":{\"info_slave_kind\":\"bond\",\"info_slave_data\":{\"state\":\"ACTIVE\",\"mii_status\":\"UP\",\"link_failure_count\":0,\"perm_hwaddr\":\"00:62:0b:e7:99:50\",\"queue_id\":0,\"prio\":0}},\"num_tx_queues\":74,\"num_rx_queues\":74,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":524280,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"phys_port_name\":\"p0\",\"phys_switch_id\":\"5099e7feff0b6200\",\"parentbus\":\"pci\",\"parentdev\":\"0000:4b:00.0\",\"vfinfo_list\":[],\"altnames\":[\"enp75s0f0np0\"],\"addr_info\":[]},{\"ifindex\":7,\"ifname\":\"ens1f1np1\",\"flags\":[\"BROADCAST\",\"MULTICAST\",\"UP\",\"LOWER_UP\"],\"mtu\":9000,\"qdisc\":\"mq\",\"operstate\":\"UP\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"ether\",\"address\":\"00:62:0b:e7:99:51\",\"broadcast\":\"ff:ff:ff:ff:ff:ff\",\"promiscuity\":0,\"allmulti\":0,\"min_mtu\":60,\"max_mtu\":9600,\"num_tx_queues\":74,\"num_rx_queues\":74,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":524280,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"phys_port_name\":\"p1\",\"phys_switch_id\":\"5099e7feff0b6200\",\"parentbus\":\"pci\",\"parentdev\":\"0000:4b:00.1\",\"vfinfo_list\":[],\"altnames\":[\"enp75s0f1np1\"],\"addr_info\":[{\"family\":\"inet\",\"local\":\"100.100.13.104\",\"prefixlen\":24,\"broadcast\":\"100.100.13.255\",\"scope\":\"global\",\"noprefixroute\":true,\"label\":\"ens1f1np1\",\"valid_life_time\":4294967295,\"preferred_life_time\":4294967295},{\"family\":\"inet6\",\"local\":\"fe80::262:bff:fee7:9951\",\"prefixlen\":64,\"scope\":\"link\",\"noprefixroute\":true,\"valid_life_time\":4294967295,\"preferred_life_time\":4294967295}]},{\"ifindex\":10,\"ifname\":\"ens3f0np0\",\"flags\":[\"NO-CARRIER\",\"BROADCAST\",\"MULTICAST\",\"UP\"],\"mtu\":1500,\"qdisc\":\"mq\",\"operstate\":\"DOWN\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"ether\",\"address\":\"00:62:0b:e6:3a:20\",\"broadcast\":\"ff:ff:ff:ff:ff:ff\",\"promiscuity\":0,\"allmulti\":0,\"min_mtu\":60,\"max_mtu\":9600,\"num_tx_queues\":74,\"num_rx_queues\":74,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":524280,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"phys_port_name\":\"p0\",\"phys_switch_id\":\"203ae6feff0b6200\",\"parentbus\":\"pci\",\"parentdev\":\"0000:b1:00.0\",\"vfinfo_list\":[],\"altnames\":[\"enp177s0f0np0\"],\"addr_info\":[]},{\"ifindex\":11,\"ifname\":\"ens3f1np1\",\"flags\":[\"NO-CARRIER\",\"BROADCAST\",\"MULTICAST\",\"UP\"],\"mtu\":1500,\"qdisc\":\"mq\",\"operstate\":\"DOWN\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"ether\",\"address\":\"00:62:0b:e6:3a:21\",\"broadcast\":\"ff:ff:ff:ff:ff:ff\",\"promiscuity\":0,\"allmulti\":0,\"min_mtu\":60,\"max_mtu\":9600,\"num_tx_queues\":74,\"num_rx_queues\":74,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":524280,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"phys_port_name\":\"p1\",\"phys_switch_id\":\"203ae6feff0b6200\",\"parentbus\":\"pci\",\"parentdev\":\"0000:b1:00.1\",\"vfinfo_list\":[],\"altnames\":[\"enp177s0f1np1\"],\"addr_info\":[]},{\"ifindex\":13,\"ifname\":\"bond0\",\"flags\":[\"BROADCAST\",\"MULTICAST\",\"MASTER\",\"UP\",\"LOWER_UP\"],\"mtu\":1500,\"qdisc\":\"noqueue\",\"master\":\"bridge0\",\"operstate\":\"UP\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"ether\",\"address\":\"00:62:0b:e7:99:50\",\"broadcast\":\"ff:ff:ff:ff:ff:ff\",\"promiscuity\":1,\"allmulti\":1,\"min_mtu\":68,\"max_mtu\":65535,\"linkinfo\":{\"info_kind\":\"bond\",\"info_data\":{\"mode\":\"active-backup\",\"active_slave\":\"ens1f0np0\",\"miimon\":100,\"updelay\":0,\"downdelay\":0,\"peer_notify_delay\":0,\"use_carrier\":1,\"arp_interval\":0,\"arp_missed_max\":2,\"arp_validate\":null,\"arp_all_targets\":\"any\",\"primary_reselect\":\"always\",\"fail_over_mac\":\"none\",\"xmit_hash_policy\":\"layer2\",\"resend_igmp\":1,\"num_peer_notif\":1,\"all_slaves_active\":0,\"min_links\":0,\"lp_interval\":1,\"packets_per_slave\":1,\"ad_lacp_active\":\"on\",\"ad_lacp_rate\":\"slow\",\"ad_select\":\"stable\",\"tlb_dynamic_lb\":1},\"info_slave_kind\":\"bridge\",\"info_slave_data\":{\"state\":\"forwarding\",\"priority\":32,\"cost\":100,\"hairpin\":false,\"guard\":false,\"root_block\":false,\"fastleave\":false,\"learning\":true,\"flood\":true,\"id\":\"0x8001\",\"no\":\"0x1\",\"designated_port\":32769,\"designated_cost\":0,\"bridge_id\":\"8000.0:62:b:e7:99:50\",\"root_id\":\"8000.0:62:b:e7:99:50\",\"hold_timer\":0.00,\"message_age_timer\":0.00,\"forward_delay_timer\":0.00,\"topology_change_ack\":0,\"config_pending\":0,\"proxy_arp\":false,\"proxy_arp_wifi\":false,\"multicast_router\":1,\"mcast_flood\":true,\"bcast_flood\":true,\"mcast_to_unicast\":false,\"neigh_suppress\":false,\"group_fwd_mask\":\"0\",\"group_fwd_mask_str\":\"0x0\",\"vlan_tunnel\":false,\"isolated\":false,\"locked\":false,\"mab\":false}},\"num_tx_queues\":16,\"num_rx_queues\":16,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":524280,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"addr_info\":[]},{\"ifindex\":14,\"ifname\":\"bridge0\",\"flags\":[\"BROADCAST\",\"MULTICAST\",\"UP\",\"LOWER_UP\"],\"mtu\":1500,\"qdisc\":\"noqueue\",\"operstate\":\"UP\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"ether\",\"address\":\"00:62:0b:e7:99:50\",\"broadcast\":\"ff:ff:ff:ff:ff:ff\",\"promiscuity\":0,\"allmulti\":0,\"min_mtu\":68,\"max_mtu\":65535,\"linkinfo\":{\"info_kind\":\"bridge\",\"info_data\":{\"forward_delay\":1500,\"hello_time\":200,\"max_age\":2000,\"ageing_time\":30000,\"stp_state\":0,\"priority\":32768,\"vlan_filtering\":0,\"vlan_protocol\":\"802.1Q\",\"bridge_id\":\"8000.0:62:b:e7:99:50\",\"root_id\":\"8000.0:62:b:e7:99:50\",\"root_port\":0,\"root_path_cost\":0,\"topology_change\":0,\"topology_change_detected\":0,\"hello_timer\":0.00,\"tcn_timer\":0.00,\"topology_change_timer\":0.00,\"gc_timer\":4.60,\"vlan_default_pvid\":1,\"vlan_stats_enabled\":0,\"vlan_stats_per_port\":0,\"group_fwd_mask\":\"0\",\"group_addr\":\"01:80:c2:00:00:00\",\"mcast_snooping\":1,\"no_linklocal_learn\":0,\"mcast_vlan_snooping\":0,\"mcast_router\":1,\"mcast_query_use_ifaddr\":0,\"mcast_querier\":0,\"mcast_hash_elasticity\":16,\"mcast_hash_max\":4096,\"mcast_last_member_cnt\":2,\"mcast_startup_query_cnt\":2,\"mcast_last_member_intvl\":100,\"mcast_membership_intvl\":26000,\"mcast_querier_intvl\":25500,\"mcast_query_intvl\":12500,\"mcast_query_response_intvl\":1000,\"mcast_startup_query_intvl\":3125,\"mcast_stats_enabled\":0,\"mcast_igmp_version\":2,\"mcast_mld_version\":1,\"nf_call_iptables\":0,\"nf_call_ip6tables\":0,\"nf_call_arptables\":0}},\"num_tx_queues\":1,\"num_rx_queues\":1,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":65536,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"addr_info\":[{\"family\":\"inet\",\"local\":\"10.10.13.4\",\"prefixlen\":16,\"broadcast\":\"10.10.255.255\",\"scope\":\"global\",\"noprefixroute\":true,\"label\":\"bridge0\",\"valid_life_time\":4294967295,\"preferred_life_time\":4294967295},{\"family\":\"inet6\",\"local\":\"fe80::d5e:1d53:eb29:3d29\",\"prefixlen\":64,\"scope\":\"link\",\"noprefixroute\":true,\"valid_life_time\":4294967295,\"preferred_life_time\":4294967295}]},{\"ifindex\":15,\"ifname\":\"vnet0\",\"flags\":[\"BROADCAST\",\"MULTICAST\",\"UP\",\"LOWER_UP\"],\"mtu\":1500,\"qdisc\":\"noqueue\",\"master\":\"bridge0\",\"operstate\":\"UNKNOWN\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"ether\",\"address\":\"fe:24:81:5c:0d:72\",\"broadcast\":\"ff:ff:ff:ff:ff:ff\",\"promiscuity\":1,\"allmulti\":1,\"min_mtu\":68,\"max_mtu\":65521,\"linkinfo\":{\"info_kind\":\"tun\",\"info_data\":{\"type\":\"tap\",\"pi\":false,\"vnet_hdr\":true,\"multi_queue\":false,\"persist\":false},\"info_slave_kind\":\"bridge\",\"info_slave_data\":{\"state\":\"forwarding\",\"priority\":32,\"cost\":100,\"hairpin\":false,\"guard\":false,\"root_block\":false,\"fastleave\":false,\"learning\":true,\"flood\":true,\"id\":\"0x8002\",\"no\":\"0x2\",\"designated_port\":32770,\"designated_cost\":0,\"bridge_id\":\"8000.0:62:b:e7:99:50\",\"root_id\":\"8000.0:62:b:e7:99:50\",\"hold_timer\":0.00,\"message_age_timer\":0.00,\"forward_delay_timer\":0.00,\"topology_change_ack\":0,\"config_pending\":0,\"proxy_arp\":false,\"proxy_arp_wifi\":false,\"multicast_router\":1,\"mcast_flood\":true,\"bcast_flood\":true,\"mcast_to_unicast\":false,\"neigh_suppress\":false,\"group_fwd_mask\":\"0\",\"group_fwd_mask_str\":\"0x0\",\"vlan_tunnel\":false,\"isolated\":false,\"locked\":false,\"mab\":false}},\"num_tx_queues\":1,\"num_rx_queues\":1,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":65536,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"addr_info\":[{\"family\":\"inet6\",\"local\":\"fe80::fc24:81ff:fe5c:d72\",\"prefixlen\":64,\"scope\":\"link\",\"valid_life_time\":4294967295,\"preferred_life_time\":4294967295}]},{\"ifindex\":17,\"ifname\":\"cloud0\",\"flags\":[\"NO-CARRIER\",\"BROADCAST\",\"MULTICAST\",\"UP\"],\"mtu\":1500,\"qdisc\":\"noqueue\",\"operstate\":\"DOWN\",\"group\":\"default\",\"txqlen\":1000,\"link_type\":\"ether\",\"address\":\"00:00:00:00:00:00\",\"broadcast\":\"ff:ff:ff:ff:ff:ff\",\"promiscuity\":0,\"allmulti\":0,\"min_mtu\":68,\"max_mtu\":65535,\"linkinfo\":{\"info_kind\":\"bridge\",\"info_data\":{\"forward_delay\":1500,\"hello_time\":200,\"max_age\":2000,\"ageing_time\":30000,\"stp_state\":0,\"priority\":32768,\"vlan_filtering\":0,\"vlan_protocol\":\"802.1Q\",\"bridge_id\":\"8000.0:0:0:0:0:0\",\"root_id\":\"8000.0:0:0:0:0:0\",\"root_port\":0,\"root_path_cost\":0,\"topology_change\":0,\"topology_change_detected\":0,\"hello_timer\":0.00,\"tcn_timer\":0.00,\"topology_change_timer\":0.00,\"gc_timer\":126.48,\"vlan_default_pvid\":1,\"vlan_stats_enabled\":0,\"vlan_stats_per_port\":0,\"group_fwd_mask\":\"0\",\"group_addr\":\"01:80:c2:00:00:00\",\"mcast_snooping\":1,\"no_linklocal_learn\":0,\"mcast_vlan_snooping\":0,\"mcast_router\":1,\"mcast_query_use_ifaddr\":0,\"mcast_querier\":0,\"mcast_hash_elasticity\":16,\"mcast_hash_max\":4096,\"mcast_last_member_cnt\":2,\"mcast_startup_query_cnt\":2,\"mcast_last_member_intvl\":100,\"mcast_membership_intvl\":26000,\"mcast_querier_intvl\":25500,\"mcast_query_intvl\":12500,\"mcast_query_response_intvl\":1000,\"mcast_startup_query_intvl\":3125,\"mcast_stats_enabled\":0,\"mcast_igmp_version\":2,\"mcast_mld_version\":1,\"nf_call_iptables\":0,\"nf_call_ip6tables\":0,\"nf_call_arptables\":0}},\"num_tx_queues\":1,\"num_rx_queues\":1,\"gso_max_size\":65536,\"gso_max_segs\":65535,\"tso_max_size\":524280,\"tso_max_segs\":65535,\"gro_max_size\":65536,\"gso_ipv4_max_size\":65536,\"gro_ipv4_max_size\":65536,\"addr_info\":[{\"family\":\"inet\",\"local\":\"169.254.0.1\",\"prefixlen\":16,\"scope\":\"global\",\"label\":\"cloud0\",\"valid_life_time\":4294967295,\"preferred_life_time\":4294967295},{\"family\":\"inet6\",\"local\":\"fe80::9405:a2ff:fe0f:7100\",\"prefixlen\":64,\"scope\":\"link\",\"valid_life_time\":4294967295,\"preferred_life_time\":4294967295}]}]")

		if err := xml.Unmarshal(stdout, &NICs); err != nil {
			utils.FancyHandleError(err)

		}
	}
	for _, NIC := range NICs {
		//if NIC.Parentbus == "pci" {
		fmt.Println(NIC.Ifname, ": ", NIC.Parentdev, "(", reflect.TypeOf(NIC.Parentdev), ")", ": ", NIC.Master)
		nic.NICs = append(nic.NICs, NIC)
		//}
	}
	//nic.NICs = NICs
	nic.RefreshTime = time.Now()
}

// Get godoc
//
//	@Summary		Show List of NIC
//	@Description	Cube의 NIC목록을 보여줍니다.
//	@Tags			API, CUBE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	TypeNICStatus
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/cube/nics [get]
func (nic *TypeNICStatus) Get(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, nic)
}
