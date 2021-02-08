package physicalnode

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

/*
相同的pod分散在不同的物理机
*/
var _ framework.FilterPlugin = &PhysicalNode{}

const (
	Name = "Physical"
)

type PhysicalNode struct {
}

func New(configuration runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	klog.V(2).Info("scheduler-plugin-custom")
	return &PhysicalNode{}, nil
}

func (p *PhysicalNode) Name() string {
	return Name
}

func (s *PhysicalNode) PreFilter(ctx context.Context, state *framework.CycleState, pod *corev1.Pod) *framework.Status {
	klog.V(3).Infof("prefilter pod: %v", pod.Name)
	klog.V(2).Infof("进入prefilter，pod: %v", pod.Name)
	return framework.NewStatus(framework.Success, "")
}
func (p *PhysicalNode) Filter(ctx context.Context, state *framework.CycleState, pod *corev1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.V(3).Infof("filter pod: %v, node: %v", pod.Name, nodeInfo.Node().Name)
	klog.V(2).Infof("进入filter，pod: %v, node: %v", pod.Name, nodeInfo.Node().Name)
	return framework.NewStatus(framework.Success, "成功剔除相同的物理机")
}

//PreBind(ctx context.Context, state *CycleState, p *v1.Pod, nodeName string) *Status
func (s *PhysicalNode) PreBind(ctx context.Context, state *framework.CycleState, pod *corev1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("PreBind pod: %v", pod.Name)
	klog.V(2).Infof("进入PreBind，pod: %v", pod.Name)
	return framework.NewStatus(framework.Success, "")
}
